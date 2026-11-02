package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppConfig      AppConfig      `envPrefix:"APP_"`
	DatabaseConfig DataBaseConfig `envPrefix:"DB_"`
	LoggerConfig   LoggerConfig   `envPrefix:"LOGGER_"`
}

type AppConfig struct {
	Env      string `env:"ENV" envDefault:"development"`
	Port     string `env:"PORT" envDefault:"8080"`
	Name     string `env:"NAME" envDefault:"rewardApp"`
	TimeZone string `env:"TIMEZONE" envDefault:"UTC"`
}

type DataBaseConfig struct {
	Host     string        `env:"HOST,required"`
	Port     string        `env:"PORT" envDefault:"5432"`
	User     string        `env:"USER,required"`
	Password string        `env:"PASSWORD,required"`
	Name     string        `env:"NAME,required"`
	SSLMode  string        `env:"SSLMODE" envDefault:"disable"`
	Timeout  time.Duration `env:"TIMEOUT" envDefault:"5s"`
}

type LoggerConfig struct {
	Level        string `env:"LEVEL" envDefault:"DEBUG"`
	LoggerFolder string `env:"FOLDER"`
}

func Load() (Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
