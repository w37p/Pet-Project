package router

import (
	"github.com/bullockz21/pet_project21/internal/bot"
	menuCtrl "github.com/bullockz21/pet_project21/internal/controller/telegram/menu" // Исправленный путь
	telegram "github.com/bullockz21/pet_project21/internal/controller/telegram/webhook"
	menuPresenter "github.com/bullockz21/pet_project21/internal/modules/presenter/menu"
	menuRepo "github.com/bullockz21/pet_project21/internal/modules/repository/menu"
	menuUsecase "github.com/bullockz21/pet_project21/internal/modules/usecase/menu"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes настраивает все маршруты Gin и возвращает роутер.
func SetupRoutes(handler *bot.Handler, db *gorm.DB) *gin.Engine { // Добавить db как параметр
	router := gin.Default()

	// Инициализация компонентов меню
	menuRepository := menuRepo.NewMenuRepository(db)
	menuUC := menuUsecase.NewMenuUseCase(menuRepository)
	menuPresenter := menuPresenter.NewMenuPresenter()
	menuController := menuCtrl.NewMenuController(menuUC, menuPresenter)

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/webhook", telegram.WebhookHandler(handler))
		apiV1.GET("/ping", telegram.PingHandler)
		apiV1.GET("/menu", menuController.GetMenu)
	}

	return router
}
