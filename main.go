package main

import (
	"docker-volume-watchdog/discord"
	"docker-volume-watchdog/discord/models"
	"docker-volume-watchdog/environment"
	"docker-volume-watchdog/watcher"
	watcherModels "docker-volume-watchdog/watcher/models"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

func main() {
	initDiscord()
	startWatchers()
}

func initDiscord() {
	discordWebhook := environment.GetDiscordWebhook()
	discord.Init(models.Config{Url: discordWebhook})
}

func startWatchers() {
	configs := buildWatcherConfigs()
	if len(configs) == 0 {
		log.Fatal("No directories mounted")
	}

	waitGroup := sync.WaitGroup{}
	for _, config := range configs {
		go func(currentConfig watcherModels.Config) {
			defer waitGroup.Done()
			watcher.Watch(currentConfig)
		}(config)

		waitGroup.Add(1)
	}

	waitGroup.Wait()
}

func buildWatcherConfigs() []watcherModels.Config {
	configs := make([]watcherModels.Config, 0)
	intervalValue := environment.GetIntervalValue()

	for _, directory := range getDirs() {
		configs = append(configs, watcherModels.Config{
			Path:          directory,
			IntervalValue: time.Duration(intervalValue) * time.Minute,
		})
	}

	return configs
}

func getDirs() []string {
	mount := environment.GetMountPoint()
	entries, err := os.ReadDir(mount)
	if err != nil {
		log.Fatal(err)
	}

	dirs := make([]string, 0)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirs = append(dirs, path.Join(mount, entry.Name()))
	}

	return dirs
}
