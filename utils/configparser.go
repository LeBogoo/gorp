package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host   string `json:"host"`
	Server string `json:"server"`
}

func (config Config) ToString() string {
	return fmt.Sprintf("{host: %s, server: %s}", config.Host, config.Server)
}

func ParseConfig() []Config {
	files, _ := os.ReadDir("config")
	configs := make([]Config, 0)

	for _, file := range files {
		var config Config

		file, _ := os.Open("config/" + file.Name())
		decoder := json.NewDecoder(file)
		decoder.Decode(&config)

		configs = append(configs, config)
	}

	return configs
}
