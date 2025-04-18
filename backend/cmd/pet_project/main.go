package main

import (
	"github.com/bullockz21/pet_project21/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	application := app.New()
	application.Run()
}
