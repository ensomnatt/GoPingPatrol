package scraper

import (
	"net/http"
	"sync"
	"time"

	"github.com/ensomnatt/gopingpatrol/checker/internal/config"
	"github.com/ensomnatt/gopingpatrol/checker/internal/logger"
	"github.com/ensomnatt/gopingpatrol/checker/internal/producer"
)

type Scraper struct {
	log  *logger.Logger
	cfg  *config.Config
	prod *producer.Producer
}

func New(log *logger.Logger, cfg *config.Config, prod *producer.Producer) *Scraper {
	return &Scraper{
		log:  log,
		cfg:  cfg,
		prod: prod,
	}
}

func (s *Scraper) Start() error {
	interval, err := time.ParseDuration(s.cfg.ScrapeInterval)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		s.log.Info("Start scraping")

		var wg sync.WaitGroup
		maxWorkers := 10
		sem := make(chan struct{}, maxWorkers)

		for _, url := range s.cfg.URLs {
			wg.Add(1)
			sem <- struct{}{}
			go func(url string) {
				defer wg.Done()
				s.checkHealth(url)
				<-sem
			}(url)
		}

		wg.Wait()
		s.log.Info("Stop scraping")

		<-ticker.C
	}
}

func (s *Scraper) checkHealth(url string) {
	resp, err := http.Get(url)
	if err != nil {
		s.log.Errorf("Error while checking health on %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.prod.Publish(url)
		s.log.Infof("New alert: %s, status code - %v", url, resp.StatusCode)
	} else {
		s.log.Infof("Checked %s", url)
	}
}
