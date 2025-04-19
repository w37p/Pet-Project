package app

import (
	"fmt"

	"github.com/bullockz21/pet_project21/internal/infrastructure/migration"
	"github.com/sirupsen/logrus"
)

// initMigrations runs database migrations with logging.
func (a *App) initMigrations() error {
	logrus.Info("🔄 Running database migrations...")
	if err := migration.Run(a.db); err != nil {
		logrus.WithError(err).Error("❌ Database migration failed")
		return fmt.Errorf("run migrations: %w", err)
	}
	logrus.Info("✅ Database migrations completed successfully")
	return nil
}
