package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type Env struct {
	PORT   string
	DB_URL string
}

func initEnv() *Env {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("viper: read config error", "error", err)
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	return &Env{PORT: port, DB_URL: dbUrl}
}
