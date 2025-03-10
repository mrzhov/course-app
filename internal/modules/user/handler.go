package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrzhov/course-app/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c echo.Context) error {
	body := new(CreateBody)

	if err := utils.ValidateBody(body, c); err != nil {
		return err
	}

	user := User{
		Email:    body.Email,
		Password: body.Password,
	}

	if err := h.service.Create(&user); err != nil {
		return utils.EchoBadRequest(err)
	}

	response := NewUserResponse(user)
	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) GetList(c echo.Context) error {
	users := new([]User)

	if err := h.service.GetList(users); err != nil {
		return utils.EchoBadRequest(err)
	}

	response := []UserResponse{}
	for _, t := range *users {
		response = append(response, NewUserResponse(t))
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) _GetById(user *User, paramId string) error {
	id := new(uint)

	if err := utils.ValidateParamId(id, paramId); err != nil {
		return err
	}

	if err := h.service.GetById(user, *id); err != nil {
		return utils.EchoBadRequest(err)
	}

	return nil
}

func (h *Handler) GetById(c echo.Context) error {
	user := new(User)

	if err := h._GetById(user, c.Param("id")); err != nil {
		return err
	}

	response := NewUserResponse(*user)
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Patch(c echo.Context) error {
	body := new(PatchBody)
	user := new(User)

	if err := utils.ValidateBody(body, c); err != nil {
		return err
	}

	if err := h._GetById(user, c.Param("id")); err != nil {
		return err
	}

	if body.Email != nil {
		user.Email = *body.Email
	}

	if err := h.service.Patch(user); err != nil {
		return utils.EchoBadRequest(err)
	}

	response := NewUserResponse(*user)
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Delete(c echo.Context) error {
	user := new(User)

	if err := h._GetById(user, c.Param("id")); err != nil {
		return err
	}

	if err := h.service.Delete(user); err != nil {
		return utils.EchoBadRequest(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetTasks(c echo.Context) error {
	user := new(User)

	if err := h._GetById(user, c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user.Tasks)
}
