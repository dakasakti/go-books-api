package config

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal()
	}

	os.Setenv("PORT", ENV.PORT)
	log.Println("Server loaded successfully ...")
}
