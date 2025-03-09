package task

import (
	"context"

	"github.com/mrzhov/course-app/internal/utils"
	"github.com/mrzhov/course-app/internal/web/tasks"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	body := request.Body

	task := Task{
		Title:       *body.Title,
		Description: *body.Description,
		Completed:   *body.Completed,
	}

	if err := h.service.Create(&task); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	response := tasks.PostTasks201JSONResponse{
		Id:          &task.ID,
		Title:       &task.Title,
		Description: &task.Description,
		Completed:   &task.Completed,
	}

	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks := new([]Task)

	if err := h.service.GetList(allTasks); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range *allTasks {
		task := tasks.Task{
			Title:       &tsk.Title,
			Description: &tsk.Description,
			Completed:   &tsk.Completed,
		}
		response = append(response, task)
	}

	return response, nil
}

// func (h *Handler) Create(c echo.Context) error {
// 	body := new(CreateBody)

// 	if err := utils.ValidateBody(body, c); err != nil {
// 		return err
// 	}

// 	task := Task{
// 		Title:       body.Title,
// 		Description: body.Description,
// 		Completed:   body.Completed,
// 	}

// 	if err := h.service.Create(&task); err != nil {
// 		return utils.EchoBadRequest(err)
// 	}

// 	return c.JSON(http.StatusCreated, task)
// }

// func (h *Handler) GetList(c echo.Context) error {
// 	tasks := new([]Task)

// 	if err := h.service.GetList(tasks); err != nil {
// 		return utils.EchoBadRequest(err)
// 	}

// 	response := []TaskResponse{}
// 	for _, t := range *tasks {
// 		response = append(response, NewTaskResponse(t))
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func (h *Handler) _GetById(task *Task, paramId string) error {
// 	id := new(uint)

// 	if err := utils.ValidateParamId(id, paramId); err != nil {
// 		return err
// 	}

// 	if err := h.service.GetById(task, *id); err != nil {
// 		return utils.EchoBadRequest(err)
// 	}

// 	return nil
// }

// func (h *Handler) GetById(c echo.Context) error {
// 	task := new(Task)

// 	if err := h._GetById(task, c.Param("id")); err != nil {
// 		return err
// 	}

// 	response := NewTaskResponse(*task)
// 	return c.JSON(http.StatusOK, response)
// }

// func (h *Handler) Patch(c echo.Context) error {
// 	body := new(PatchBody)
// 	task := new(Task)

// 	if err := utils.ValidateBody(body, c); err != nil {
// 		return err
// 	}

// 	if err := h._GetById(task, c.Param("id")); err != nil {
// 		return err
// 	}

// 	if body.Title != nil {
// 		task.Title = *body.Title
// 	}

// 	if body.Description != nil {
// 		task.Description = *body.Description
// 	}

// 	if body.Completed != nil {
// 		task.Completed = *body.Completed
// 	}

// 	if err := h.service.Patch(task); err != nil {
// 		return utils.EchoBadRequest(err)
// 	}

// 	return c.JSON(http.StatusOK, *task)
// }

// func (h *Handler) Delete(c echo.Context) error {
// 	task := new(Task)

// 	if err := h._GetById(task, c.Param("id")); err != nil {
// 		return err
// 	}

// 	if err := h.service.Delete(task); err != nil {
// 		return utils.EchoBadRequest(err)
// 	}

// 	return c.String(http.StatusOK, "Deleted successfully!")
// }
