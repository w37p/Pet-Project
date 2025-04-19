package main

import (
	"os"

	"github.com/bullockz21/pet_project21/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Настройка логгера
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Перехват неожиданных ошибок
	defer func() {
		if r := recover(); r != nil {
			logrus.WithField("module", "main").
				WithField("error", r).
				Fatal("Unhandled panic occurred, shutting down the application")
			os.Exit(1)
		}
	}()

	logrus.Info("Setting Gin to Release Mode")
	gin.SetMode(gin.ReleaseMode)

	logrus.WithField("module", "main").Info("Creating new application instance")
	application := app.New()

	logrus.WithField("module", "main").Info("Running application")
	application.Run()
}
