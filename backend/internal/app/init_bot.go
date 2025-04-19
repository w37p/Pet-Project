package app

import (
	"fmt"

	botPkg "github.com/bullockz21/pet_project21/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

// initBot встраивает Telegram Bot и настраивает Webhook.
func (a *App) initBot() {
	// Инициализация бота
	bot, err := botPkg.NewBot(a.cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize Telegram bot")
	}
	a.bot = bot
	logrus.Infof("Telegram bot initialized: @%s", bot.Self.UserName)

	// Настройка вебхука
	a.setupWebhook()
}

// setupWebhook регистрирует Webhook URL для приема обновлений.
func (a *App) setupWebhook() {
	webhookURL := fmt.Sprintf("%s/api/v1/webhook", a.cfg.Telegram.WebhookURL)
	logrus.Infof("Configuring Telegram webhook at %s", webhookURL)

	hook, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create Telegram webhook configuration")
	}

	_, err = a.bot.Request(hook)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to set Telegram webhook")
	}
	logrus.Info("Telegram webhook set successfully")
}
