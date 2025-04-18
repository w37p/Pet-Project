package app

import (
	"github.com/bullockz21/pet_project21/configs"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type App struct {
	cfg    *configs.Config
	db     *gorm.DB
	bot    *tgbotapi.BotAPI
	router *gin.Engine
}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	a.initConfig()
	a.initDB()
	a.initBot()
	a.initServer()
	a.start()
}
