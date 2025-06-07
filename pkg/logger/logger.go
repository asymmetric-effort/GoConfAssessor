// file: logger/logger.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger is the global logrus.Logger instance.
var Logger = logrus.New()

const (
	DebugStr = "debug"
	InfoStr  = "info"
)

func init() {
	// Set the format (text, JSON, etc.) and default level here.
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetLevel(logrus.InfoLevel)
}

// SetLevel changes the global log level at runtime.
// Valid values: "panic", "fatal", "error", "warn", "info", "debug", "trace".
func SetLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	Logger.SetLevel(lvl)
	return nil
}

func IsDebug(debug bool) string {
	if debug {
		return DebugStr
	} else {
		return InfoStr
	}
}
