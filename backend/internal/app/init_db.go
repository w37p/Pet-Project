package app

import (
	"github.com/bullockz21/pet_project21/internal/infrastructure/database"
	"github.com/bullockz21/pet_project21/internal/infrastructure/migration"
	"github.com/sirupsen/logrus"
)

// initDB устанавливает соединение с базой данных и запускает миграции.
func (a *App) initDB() {
	logrus.WithField("module", "database").Info("Connecting to PostgreSQL database")
	db, err := database.NewPostgresDB(a.cfg)
	if err != nil {
		logrus.WithField("module", "database").WithError(err).Fatal("Failed to connect to database")
	}
	a.db = db

	logrus.WithField("module", "database").Info("Running database migrations")
	if err := migration.Run(db); err != nil {
		logrus.WithField("module", "database").WithError(err).Fatal("Database migration failed")
	}
	logrus.WithField("module", "database").Info("Database initialized and migrations applied successfully")
}

// closeDB закрывает соединение с базой данных.
func (a *App) closeDB() {
	logrus.WithField("module", "database").Info("Closing database connection")
	if err := database.Close(a.db); err != nil {
		logrus.WithField("module", "database").WithError(err).Error("Error closing database connection")
		return
	}
	logrus.WithField("module", "database").Info("Database connection closed successfully")
}
