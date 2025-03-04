package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	DB "github.com/mrzhov/course-app/internal/common/db"
	"github.com/mrzhov/course-app/internal/common/routes"
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

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	env := initEnv()
	db := DB.Init(env.dbUrl)
	e := initEcho()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.Register(e, db)

	if err := e.Start(env.port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
