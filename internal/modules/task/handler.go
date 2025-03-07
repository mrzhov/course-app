package task

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

	task := Task{
		Title:       body.Title,
		Description: body.Description,
		Completed:   body.Completed,
	}

	if err := h.service.Create(&task); err != nil {
		return utils.EchoBadRequest(err)
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetList(c echo.Context) error {
	tasks := new([]Task)

	if err := h.service.GetList(tasks); err != nil {
		return utils.EchoBadRequest(err)
	}

	response := []TaskResponse{}
	for _, t := range *tasks {
		response = append(response, NewTaskResponse(t))
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetById(c echo.Context) error {
	id := new(uint)
	task := new(Task)

	if err := utils.ValidateParamId(id, c.Param("id")); err != nil {
		return err
	}

	if err := h.service.GetById(task, *id); err != nil {
		return utils.EchoBadRequest(err)
	}

	response := NewTaskResponse(*task)
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Patch(c echo.Context) error {
	return nil
}

func (h *Handler) Delete(c echo.Context) error {
	return nil
}
