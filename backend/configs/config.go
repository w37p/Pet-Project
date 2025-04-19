package configs

import (
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	HTTPPort     string        `env:"HTTP_PORT,required"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

type TelegramConfig struct {
	Token      string `env:"TELEGRAM_BOT_TOKEN,required"`
	WebhookURL string `env:"WEBHOOK_URL,required"`
}

type Config struct {
	Server   ServerConfig
	Database DBConfig
	Telegram TelegramConfig
}

func Load() (*Config, error) {
	log := logrus.WithField("module", "config")

	log.Info("Loading .env file...")
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Warn(".env file not found or failed to load, using environment variables")
	} else {
		log.Info(".env file loaded successfully")
	}

	var cfg Config

	log.Info("Parsing server config...")
	if err := env.Parse(&cfg.Server); err != nil {
		log.WithError(err).Error("Failed to parse server config")
		return nil, fmt.Errorf("load server config: %w", err)
	}
	log.Info("Server config loaded")

	log.Info("Parsing database config...")
	if err := env.Parse(&cfg.Database); err != nil {
		log.WithError(err).Error("Failed to parse database config")
		return nil, fmt.Errorf("load database config: %w", err)
	}
	log.Info("Database config loaded")

	log.Info("Parsing telegram config...")
	if err := env.Parse(&cfg.Telegram); err != nil {
		log.WithError(err).Error("Failed to parse telegram config")
		return nil, fmt.Errorf("load telegram config: %w", err)
	}
	log.Info("Telegram config loaded")

	if !strings.HasPrefix(cfg.Telegram.WebhookURL, "https://") {
		log.WithField("webhook_url", cfg.Telegram.WebhookURL).
			Error("WEBHOOK_URL must start with https://")
		return nil, fmt.Errorf("WEBHOOK_URL must start with https://")
	}

	log.Info("All configs loaded successfully")
	return &cfg, nil
}
