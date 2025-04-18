package main

import (
	"github.com/bullockz21/pet_project21/internal/app"
	"github.com/gin-gonic/gin"
)

// @title Beer Bot API
// @version 1.0
// @description Backend API for Telegram Beer Bot
// @contact.name Кирилл
// @contact.email kirill@example.com
// @host localhost:8080
// @BasePath /api/v1
func main() {
	gin.SetMode(gin.ReleaseMode)

	application := app.New()
	application.Run()
}
