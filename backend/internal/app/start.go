package app

import (
	"log"
)

func (a *App) start() {
	defer a.closeDB()

	log.Printf("Бот запущен: %s", a.bot.Self.UserName)

	if err := a.router.Run(":" + a.cfg.Server.HTTPPort); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
