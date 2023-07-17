package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser      string `env:"ORACLE_DB_USER"`
	DBPassword  string `env:"ORACLE_DB_PASSWORD"`
	DBAddress   string `env:"ORACLE_DB_ADDRESS"`
	DBPort      int    `env:"ORACLE_DB_PORT"`
	DBService   string `env:"ORACLE_DB_SERVICE"`
	SmtpAddress string `env:"SMPT_ADDRESS"`
	SmtpPort    int    `env:"SMPT_PORT"`
	FromName    string `env:"FROM_NAME"`
	FromAddress string `env:"FROM_ADDRESS"`
	ToAdress    string `env:"TO_ADDRESS"`
}

func New() (*Config, error) {
	// .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
