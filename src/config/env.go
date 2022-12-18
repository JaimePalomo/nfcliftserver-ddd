package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"log"
)

type config struct {
	Environment string `env:"ENV"`

	ApiPort string `env:"API_PORT"`

	Mysql_nfclift_username string `env:"MYSQL_USERS_USERNAME"`
	Mysql_nfclift_password string `env:"MYSQL_USERS_PASSWORD"`
	Mysql_nfclift_host     string `env:"MYSQL_USERS_HOST"`
	Mysql_nfclift_schema   string `env:"MYSQL_USERS_SCHEMA"`
}

var (
	Cfg config
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(&Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
