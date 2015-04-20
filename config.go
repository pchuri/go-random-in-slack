package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	WebhookUrl string `json:"webhook_url"`
	Channel    string `json:"channel"`
	Username   string `json:"username"`
}

func ReadConfig() (*Config, error) {
	for _, path := range []string{"./slack.conf"} {
		file, err := os.Open(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}

		json.NewDecoder(file)
		conf := Config{}
		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			return nil, err
		}
		return &conf, nil
	}

	return nil, errors.New("Config file not found")
}
