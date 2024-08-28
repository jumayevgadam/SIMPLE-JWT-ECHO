package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres Postgres `env:"connString"`
	Server   struct {
		PORT string `env:"PORT"`
	}
}

// Postgres is
type Postgres struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Name     string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
	SslMode  string `env:"DB_SSLMODE"`
}

// LoadConfig is
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("[config][LoadConfig]: %w", err)
	}

	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig.Process: %v", err.Error())
	}

	port, exists := os.LookupEnv("PORT")
	if !exists {
		return nil, fmt.Errorf("LoadConfig.PORT")
	}
	c.Server.PORT = port

	// Validate the config
	err = validator.New().Struct(c)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig.Validate: %v", err.Error())
	}

	return &c, nil
}
