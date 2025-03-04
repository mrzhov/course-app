package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	DB "github.com/mrzhov/course-app/internal/common/db"
	"github.com/mrzhov/course-app/internal/task"
)

type Env struct {
	port  string
	dbUrl string
}

func initEnv() Env {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	return Env{port, dbUrl}
}

func initEcho() *echo.Echo {
	e := echo.New()

	e.Group("/api")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func main() {
	env := initEnv()

	e := initEcho()
	db := DB.Init(env.dbUrl)

	task.RegisterRoutes(e, db)

	if err := e.Start(env.port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
