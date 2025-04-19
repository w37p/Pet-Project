package bot

import (
	"github.com/bullockz21/pet_project21/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func NewBot(cfg *configs.Config) (*tgbotapi.BotAPI, error) {
	// Настроим логгер logrus
	log := logrus.New()

	// Добавим информацию о модуле в логи
	log.WithFields(logrus.Fields{
		"module": "bot",
		"token":  cfg.Telegram.Token,
	}).Info("Initializing Telegram bot")

	// Создаем нового бота
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to initialize bot")
		return nil, err
	}

	// Логируем успешную инициализацию бота
	log.WithFields(logrus.Fields{
		"bot_name": bot.Self.UserName,
	}).Info("Bot successfully initialized")

	return bot, nil
}
