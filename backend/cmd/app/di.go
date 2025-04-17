package app

import (
	"fmt"
	"log"

	"github.com/bullockz21/pet_project21/configs"
	botPkg "github.com/bullockz21/pet_project21/internal/bot"
	telegramCtrl "github.com/bullockz21/pet_project21/internal/controller/telegram"
	dbpkg "github.com/bullockz21/pet_project21/internal/infrastructure/database"
	"github.com/bullockz21/pet_project21/internal/infrastructure/migration"
	userPresPkg "github.com/bullockz21/pet_project21/internal/modules/presenter/user"
	userRepoPkg "github.com/bullockz21/pet_project21/internal/modules/repository/user"
	userUCpkg "github.com/bullockz21/pet_project21/internal/modules/usecase/user"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New собирает все зависимости и возвращает готовый App.
func New() (*App, error) {
	// 1. Конфиг
	cfg, err := configs.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	// 2. Подключение к БД
	db, err := dbpkg.NewPostgresDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	// 3. Миграции
	if err := migration.Run(db); err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	// 4. Инициализация бота
	bot, err := botPkg.NewBot(cfg)
	if err != nil {
		return nil, fmt.Errorf("init bot: %w", err)
	}

	webhookURL := cfg.Telegram.WebhookURL + "/api/v1/webhook"
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatalf("Ошибка создания вебхука: %v", err)
	}
	if _, err = bot.Request(webhookConfig); err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}

	// 5. DI для модуля user
	userRepo := userRepoPkg.NewUserRepository(db)
	userUC := userUCpkg.NewUserUseCase(userRepo)
	userPres := userPresPkg.NewUserPresenter(bot)

	// 6. Контроллеры телеграма
	startH := telegramCtrl.NewStartHandler(userUC, userPres, cfg)
	cmdH := telegramCtrl.NewCommandHandler(startH, userPres)
	cbH := telegramCtrl.NewCallbackHandler(bot)
	handler := telegramCtrl.NewHandler(bot, cmdH, cbH)

	// 7. Gin + роуты + swagger
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	// Регистрация маршрутов v1
	v1 := engine.Group("/api/v1")
	telegramCtrl.RegisterWebhookRoutes(v1, handler)
	telegramCtrl.RegisterPingRoutes(v1)
	// TODO: RegisterOrderRoutes(v1), RegisterUserRoutes(v1), …

	// Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &App{
		cfg:    cfg,
		engine: engine,
		bot:    bot,
	}, nil
}
