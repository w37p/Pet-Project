package app

import (
	_ "github.com/bullockz21/pet_project21/docs"
	"github.com/bullockz21/pet_project21/internal/bot"

	//user
	telegramController "github.com/bullockz21/pet_project21/internal/bot"
	userPresenterPkg "github.com/bullockz21/pet_project21/internal/modules/presenter/user"
	userRepositoryPkg "github.com/bullockz21/pet_project21/internal/modules/repository/user"
	userUsecasePkg "github.com/bullockz21/pet_project21/internal/modules/usecase/user"

	//router
	router "github.com/bullockz21/pet_project21/internal/router/v1"

	//menu
	//menuCtrl "github.com/bullockz21/pet_project21/internal/controller/telegram/menu"
	menuPresenter "github.com/bullockz21/pet_project21/internal/modules/presenter/menu"
	menuRepo "github.com/bullockz21/pet_project21/internal/modules/repository/menu"
	menuUsecase "github.com/bullockz21/pet_project21/internal/modules/usecase/menu"

	//swagger
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// initServer конфигурирует HTTP-сервер, регистрирует маршруты и Swagger.
func (a *App) initServer() {
	logrus.Info("Initializing HTTP server and routes...")
	handler := a.initHandlers()

	logrus.Info("Setting up router with registered handlers...")
	a.router = router.SetupRoutes(handler, a.db)

	logrus.Info("Adding Swagger UI endpoint at /swagger/*any...")
	a.addSwagger()
}

// initHandlers создаёт и возвращает корневой Telegram-обработчик для всех входящих обновлений.
func (a *App) initHandlers() *bot.Handler {
	logrus.Info("Initializing Telegram handlers...")

	userRepo := userRepositoryPkg.NewUserRepository(a.db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(a.bot)

	menuRep := menuRepo.NewMenuRepository(a.db)
	menuUC := menuUsecase.NewMenuUseCase(menuRep)
	menuPresenter := menuPresenter.NewMenuPresenter()

	startHandler := telegramController.NewStartHandler(userUC, userPresenter, a.cfg)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter)
	callbackHandler := telegramController.NewCallbackHandler(a.bot)

	handler := telegramController.NewHandler(a.bot, commandHandler, callbackHandler)
	logrus.Infof("Handlers initialized: %T, %T, %T", startHandler, commandHandler, callbackHandler)

	return handler
}

// addSwagger регистрирует эндпоинт для Swagger UI.
func (a *App) addSwagger() {
	logrus.Info("Registering Swagger UI endpoints...")
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
