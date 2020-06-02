package logger

import (
	"testing"
)

func TestNewLoggerWithoutSyslogEndpoint(t *testing.T) {
	logger := NewLogger("my-package", "DEBUG", nil)
	if logger == nil {
		t.FailNow()
	}
	logger.Debugf("success")
}

func TestNewLoggerWithSyslogEndpoint(t *testing.T) {
	endpoint := "localhost:514"
	logger := NewLogger("my-package", "INFO", &endpoint)
	if logger == nil {
		t.FailNow()
	}
	logger.Debugf("success")
}
