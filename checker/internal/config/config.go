package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	ScrapeInterval string   `toml:"scrape_interval"`
	URLs           []string `toml:"urls"`
}

func Load() (*Config, error) {
	doc, err := os.ReadFile("/app/config.toml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = toml.Unmarshal(doc, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
