package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email       string `env:"EMAIL"`
	Password    string `env:"PASSWORD"`
	Address     string `env:"ADDRESS"`
	StoragePath string `env:"STORAGE_PATH"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("LoadConfig error: %w", err)
	}

	cfg := &Config{
		Email:       os.Getenv("EMAIL"),
		Password:    os.Getenv("PASSWORD"),
		Address:     os.Getenv("ADDRESS"),
		StoragePath: os.Getenv("STORAGE_PATH"),
	}

	return cfg, nil
}
