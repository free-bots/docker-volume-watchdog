package discord

import (
	"bytes"
	"docker-volume-watchdog/discord/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var currentConfig *models.Config

func Init(config models.Config) {
	currentConfig = &config
}

func Notify(message string) error {
	if currentConfig == nil {
		return fmt.Errorf("no discord config provided")
	}

	body := models.Message{Content: message}

	marshaledJson, err := json.Marshal(body)

	if err != nil {
		return err
	}

	jsonReader := bytes.NewReader(marshaledJson)

	response, err := http.Post(currentConfig.Url, "application/json", jsonReader)

	if err != nil {
		return err
	}

	fmt.Printf("Send message to discord with status %s\n", response.Status)

	return nil
}
