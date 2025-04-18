package app

import (
	"github.com/bullockz21/pet_project21/internal/bot"
	"github.com/bullockz21/pet_project21/internal/router/v1"

	_ "github.com/bullockz21/pet_project21/docs"
	telegramController "github.com/bullockz21/pet_project21/internal/bot"
	userPresenterPkg "github.com/bullockz21/pet_project21/internal/modules/presenter/user"
	userRepositoryPkg "github.com/bullockz21/pet_project21/internal/modules/repository/user"
	userUsecasePkg "github.com/bullockz21/pet_project21/internal/modules/usecase/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *App) initServer() {
	handler := a.initHandlers()
	a.router = router.SetupRoutes(handler)
	a.addSwagger()
}

func (a *App) initHandlers() *bot.Handler {
	userRepo := userRepositoryPkg.NewUserRepository(a.db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(a.bot)
	startHandler := telegramController.NewStartHandler(userUC, userPresenter, a.cfg)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter)
	callbackHandler := telegramController.NewCallbackHandler(a.bot)

	return telegramController.NewHandler(a.bot, commandHandler, callbackHandler)
}

func (a *App) addSwagger() {
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
