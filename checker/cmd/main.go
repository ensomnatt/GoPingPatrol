package main

import (
	"github.com/ensomnatt/gopingpatrol/checker/internal/config"
	"github.com/ensomnatt/gopingpatrol/checker/internal/logger"
	"github.com/ensomnatt/gopingpatrol/checker/internal/scraper"
)

func main() {
	log := logger.New("debug")
	cfg := config.Load(log)
	scraper := scraper.New(log, cfg)

	scraper.Start()
}
