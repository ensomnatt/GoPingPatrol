package main

import (
	"github.com/ensomnatt/gopingpatrol/checker/internal/config"
	"github.com/ensomnatt/gopingpatrol/checker/internal/logger"
	"github.com/ensomnatt/gopingpatrol/checker/internal/producer"
	"github.com/ensomnatt/gopingpatrol/checker/internal/scraper"
)

func main() {
	log := logger.New("debug")
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	prod, err := producer.New(log)
	if err != nil {
		log.Fatalf("Failed to connect to rabbitmq: %v", err)
	}
	scraper := scraper.New(log, cfg, prod)

	err = scraper.Start()
	if err != nil {
		log.Fatalf("Failed to start scraper: %v", err)
	}
}
