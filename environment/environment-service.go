package environment

import (
	"log"
	"os"
	"strconv"
)

const envPrefix = "DOCKER_VOLUME_WATCHDOG"

func GetDiscordWebhook() string {
	discordWebhook, ok := os.LookupEnv(envPrefix + "_DISCORD_WEBHOOK")
	if !ok {
		log.Fatal("No discord webhook url provided")
	}

	return discordWebhook
}

func GetIntervalValue() int {
	intervalValue, ok := os.LookupEnv(envPrefix + "_INTERVAL_VALUE")
	if !ok {
		log.Fatal("Interval value not defined")
	}

	intervalValueAsInt, err := strconv.ParseInt(intervalValue, 10, 0)
	if err != nil {
		log.Fatal("Could not parse interval value")
	}

	return int(intervalValueAsInt)
}

func GetMountPoint() string {
	return "/watch-dog-mount"
}
