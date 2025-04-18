package app

import (
	"log"

	"github.com/bullockz21/pet_project21/internal/infrastructure/database"
	"github.com/bullockz21/pet_project21/internal/infrastructure/migration"
)

func (a *App) initDB() {
	db, err := database.NewPostgresDB(a.cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	a.db = db

	if err := migration.Run(db); err != nil {
		log.Fatalf("Миграция не удалась: %v", err)
	}
}

func (a *App) closeDB() {
	if err := database.Close(a.db); err != nil {
		log.Printf("Ошибка закрытия базы данных: %v", err)
	}
}
