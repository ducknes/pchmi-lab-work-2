package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configPath = "config/config.json"
)

type Config struct {
	Sqlite         string `json:"sqlite"`
	Postgres       string `json:"postgres"`
	DatabaseDriver string `json:"driver"`
}

func New() (Config, error) {
	var cfg Config

	content, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("не удалость прочитать конфиг из json файла. ошибка: %w", err)
	}

	if err := json.Unmarshal(content, &cfg); err != nil {
		return Config{}, fmt.Errorf("не удалость распарсить конфиг из json файла в структуру. ошибка: %w", err)
	}

	return cfg, nil
}
