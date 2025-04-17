package configs

import (
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
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
	_ = godotenv.Load("../../configs/.env")

	var cfg Config
	if err := env.Parse(&cfg.Server); err != nil {
		return nil, fmt.Errorf("load server config: %w", err)
	}
	if err := env.Parse(&cfg.Database); err != nil {
		return nil, fmt.Errorf("load database config: %w", err)
	}
	if err := env.Parse(&cfg.Telegram); err != nil {
		return nil, fmt.Errorf("load telegram config: %w", err)
	}

	if !strings.HasPrefix(cfg.Telegram.WebhookURL, "https://") {
		return nil, fmt.Errorf("WEBHOOK_URL must start with https://")
	}

	return &cfg, nil
}
