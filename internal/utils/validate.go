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

func ValidateParamId(id *uint, paramId string) error {
	if err := validate.Var(paramId, "required,number"); err != nil {
		return EchoBadRequest(err)
	}

	idInt, idIntErr := strconv.Atoi(paramId)
	if idIntErr != nil {
		return EchoBadRequest(idIntErr)
	}

	*id = uint(idInt)

	return nil
}
