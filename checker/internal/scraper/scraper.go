package scraper

import (
	"net/http"
	"sync"

	"github.com/ensomnatt/gopingpatrol/checker/internal/config"
	"github.com/ensomnatt/gopingpatrol/checker/internal/logger"
)

type Scraper struct {
	log *logger.Logger
	cfg config.Config
}

func New(log *logger.Logger, cfg config.Config) *Scraper {
	return &Scraper{
		log: log,
		cfg: cfg,
	}
}

func (s *Scraper) Start() {
	var wg sync.WaitGroup
	maxWorkers := 10
	sem := make(chan struct{}, maxWorkers)

	s.log.Info("Start scraping")
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
}

func (s *Scraper) checkHealth(url string) {
	resp, err := http.Get(url)
	if err != nil {
		s.log.Errorf("Error while checking health on %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.log.Infof("New alert: %s, status code - %v", url, resp.StatusCode)
	} else {
		s.log.Infof("Checked %s", url)
	}
}
