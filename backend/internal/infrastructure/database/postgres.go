package database

import (
	"fmt"

	"github.com/bullockz21/pet_project21/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *configs.Config) (*gorm.DB, error) {
	log := logrus.WithField("module", "database")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	log.Info("Connecting to the PostgreSQL database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Error("Failed to connect to the PostgreSQL database")
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	log.Info("Successfully connected to the PostgreSQL database")
	return db, nil
}

func Close(db *gorm.DB) error {
	log := logrus.WithField("module", "database")

	sqlDB, err := db.DB()
	if err != nil {
		log.WithError(err).Error("Failed to get raw database instance")
		return fmt.Errorf("get raw DB error: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.WithError(err).Error("Failed to close database connection")
		return err
	}

	log.Info("Database connection closed successfully")
	return nil
}
