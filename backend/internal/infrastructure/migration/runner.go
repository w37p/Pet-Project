package migration

import (
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	log := logrus.WithField("module", "migration")

	sqlDB, err := db.DB()
	if err != nil {
		log.WithError(err).Error("Failed to get sql.DB from gorm.DB")
		return fmt.Errorf("get sql.DB failed: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.WithError(err).Error("Failed to set goose dialect to postgres")
		return err
	}

	const migrationDir = "../../migrations"
	log.Infof("Running migrations from directory: %s", migrationDir)

	if err := goose.Up(sqlDB, migrationDir); err != nil {
		log.WithError(err).Error("Migration failed")
		return fmt.Errorf("миграция не удалась: %v", err)
	}

	log.Info("Migrations applied successfully")
	return nil
}
