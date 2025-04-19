package app

import (
	"github.com/bullockz21/pet_project21/configs"
	"github.com/sirupsen/logrus"
)

func (a *App) initConfig() {
	// Логируем начало загрузки конфигурации
	logrus.WithField("module", "config").Info("Start loading configuration")

	cfg, err := configs.Load()
	if err != nil {
		// Логируем ошибку и прерываем выполнение
		logrus.
			WithField("module", "config").
			WithError(err).
			Fatal("Failed to load configuration")
	}

	// При желании можно вывести ключевые поля конфига (не выводите чувствительные данные!)
	logrus.
		WithFields(logrus.Fields{
			"module":      "config",
			"http_port":   cfg.Server.HTTPPort,
			"db_host":     cfg.Database.Host,
			"db_name":     cfg.Database.Name,
			"webhook_url": cfg.Telegram.WebhookURL,
		}).
		Info("Configuration loaded successfully")

	a.cfg = cfg
}
