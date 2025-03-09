package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrzhov/course-app/internal/utils"
	"github.com/rs/zerolog"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return utils.EchoBadRequest(err)
	}
	return nil
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Group("/api")

	// e.Use(middleware.Logger())
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}
