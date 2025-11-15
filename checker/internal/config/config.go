package config

import (
	"os"

	"github.com/ensomnatt/gopingpatrol/checker/internal/logger"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	ScrapeInterval string   `toml:"scrape_interval"`
	URLs           []string `toml:"urls"`
}

func Load(log *logger.Logger) Config {
	doc, err := os.ReadFile("/app/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = toml.Unmarshal(doc, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
