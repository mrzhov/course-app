package config

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Config struct {
	Env  *Env
	Echo *echo.Echo
	DB   *gorm.DB
}

func Init() *Config {
	env := initEnv()
	echo := initEcho()
	db := initDb(env.DB_URL)

	return &Config{Env: env, Echo: echo, DB: db}
}
