package watcher

import (
	"docker-volume-watchdog/discord"
	"docker-volume-watchdog/watcher/models"
	"fmt"
	"os"
	"time"
)

const watchFile = ".docker-volume-watchdog"

func Watch(config models.Config) {
	fmt.Printf("Start watching for %s\n", config.Path)
	for true {
		timer := time.NewTimer(config.IntervalValue)
		<-timer.C

		if directoryExists(config.Path) {
			continue
		}

		err := discord.Notify(fmt.Sprintf("Directory %s not found", config.Path))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func directoryExists(path string) bool {
	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return false
	}

	for _, entry := range entries {
		if entry.Name() == watchFile {
			return true
		}
	}

	return false
}
