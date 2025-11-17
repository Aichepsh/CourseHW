package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    Email
	Password Password
	Address  Address
}
type Email struct {
	Email string
}
type Password struct {
	Password string
}

type Address struct {
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &Config{
		Email: Email{
			Email: os.Getenv("EMAIL"),
		},
		Password: Password{
			Password: os.Getenv("PASSWORD"),
		},
		Address: Address{
			Address: os.Getenv("ADDRESS"),
		},
	}

}
