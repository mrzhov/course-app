package config

import "github.com/spf13/viper"

type Env struct {
	PORT   string
	DB_URL string
}

func initEnv() *Env {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	return &Env{PORT: port, DB_URL: dbUrl}
}
