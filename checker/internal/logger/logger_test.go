package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  logrus.Level
	}{
		{"debug", "debug", logrus.DebugLevel},
		{"warn", "warn", logrus.WarnLevel},
		{"error", "error", logrus.ErrorLevel},
		{"default", "something", logrus.InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := New(tt.level)

			if log.Level != tt.want {
				t.Errorf("Expected level %v, got %v", tt.want, log.Level)
			}

			if log.Formatter == nil {
				t.Errorf("Expected formatter to be set")
			}
		})
	}
}
