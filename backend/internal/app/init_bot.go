package app

import (
	"log"

	botPkg "github.com/bullockz21/pet_project21/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *App) initBot() {
	bot, err := botPkg.NewBot(a.cfg)
	if err != nil {
		log.Fatalf("Bot init failed: %v", err)
	}
	a.bot = bot

	a.setupWebhook()
}

func (a *App) setupWebhook() {
	webhookURL := a.cfg.Telegram.WebhookURL + "/api/v1/webhook"
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatalf("Ошибка создания вебхука: %v", err)
	}

	if _, err = a.bot.Request(webhookConfig); err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}
}
