package app

import (
	"github.com/bullockz21/pet_project21/configs"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	cfg    *configs.Config
	db     *gorm.DB
	bot    *tgbotapi.BotAPI
	router *gin.Engine
}

func New() *App {
	logrus.WithField("module", "app").Info("Creating new application instance")
	return &App{}
}

func (a *App) Run() {
	log := logrus.WithField("module", "app")

	log.Info("Initializing configuration...")
	a.initConfig()

	log.Info("Initializing database...")
	a.initDB()

	log.Info("Initializing Telegram bot...")
	a.initBot()

	log.Info("Initializing HTTP server...")
	a.initServer()

	log.Info("Starting application...")
	a.start()
}
