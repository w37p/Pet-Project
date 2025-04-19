package app

import (
	"github.com/sirupsen/logrus"
)

func (a *App) start() {
	defer a.closeDB()

	log := logrus.WithField("module", "app")

	log.Infof("🚀 Бот запущен: %s", a.bot.Self.UserName)

	if err := a.router.Run(":" + a.cfg.Server.HTTPPort); err != nil {
		log.WithError(err).Fatal("❌ Ошибка запуска HTTP сервера")
	}
}
