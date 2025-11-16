package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config-*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := []byte(`
		scrape_interval = "10s"
		URLs = ["http://youtube.com/health", "http://twitch.tv/health"]
	`)

	_, err = tmpFile.Write(content)
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	ConfigPath = tmpFile.Name()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.ScrapeInterval != "10s" {
		t.Errorf("Got %s, want 10s", cfg.ScrapeInterval)
	}

	if len(cfg.URLs) != 2 {
		t.Errorf("Expected 2 urls, got %d", len(cfg.URLs))
	}

	if cfg.URLs[0] != "http://youtube.com/health" {
		t.Errorf("First url mismatch")
	}
}
