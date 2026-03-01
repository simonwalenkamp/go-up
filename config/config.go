package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Urls []string `json:"urls"`
}

func Load() Config {
	f, _ := os.Open("config.json")
	var config Config
	json.NewDecoder(f).Decode(&config)

	return config
}
