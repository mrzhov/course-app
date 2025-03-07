package utils

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateBody(body any, c echo.Context) error {
	if err := c.Bind(body); err != nil {
		return EchoBadRequest(err)
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	return nil
}

func ValidateId(id string) (uint, error) {
	if err := validate.Var(id, "required,number"); err != nil {
		return 0, EchoBadRequest(err)
	}

	idInt, idIntErr := strconv.Atoi(id)
	if idIntErr != nil {
		return 0, EchoBadRequest(idIntErr)
	}

	return uint(idInt), nil
}
