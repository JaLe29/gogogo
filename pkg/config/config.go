package internal

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func Load() (config Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file. Skipping.")
	}

	err = env.Parse(&config)
	return
}

type Config struct {
	CokieTTL    int    `env:"COOKIE_TTL,required"`
	DatabaseUrl string `env:"DATABASE_URL,required"`
	AdminDomain string `env:"ADMIN_DOMAIN,required"`
}
