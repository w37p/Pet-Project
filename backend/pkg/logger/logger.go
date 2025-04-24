// internal/logger/logger.go
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(environment string) {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)

	switch environment {
	case "production":
		Log.SetLevel(logrus.InfoLevel)
		Log.SetFormatter(&logrus.JSONFormatter{})
	case "development":
		fallthrough
	default:
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
	}
}

func WithComponent(component string) *logrus.Entry {
	return Log.WithField("component", component)
}
