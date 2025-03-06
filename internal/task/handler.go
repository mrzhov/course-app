package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrzhov/course-app/internal/common/utils"
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetList(c echo.Context) error {
	tasks := new([]Task)

	if err := h.service.GetList(tasks); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := []TaskResponse{}

	for _, t := range *tasks {
		item := TaskResponse{
			Id:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
		}
		response = append(response, item)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetById(c echo.Context) error {
	task := new(Task)
	id, idErr := utils.ValidateId(c.Param("id"))

	if idErr != nil {
		return idErr
	}

	if err := h.service.GetById(task, id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := TaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Patch(c echo.Context) error {
	return nil
}

func (h *Handler) Delete(c echo.Context) error {
	return nil
}
