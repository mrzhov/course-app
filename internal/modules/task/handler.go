package task

import (
	"context"
	"errors"
	"strconv"

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

	if body.Title == nil {
		return nil, utils.EchoBadRequest(errors.New("field title is required"))
	}

	task := Task{
		Title: *body.Title,
	}

	if body.Description != nil {
		task.Description = *body.Description
	}

	if body.Completed != nil {
		task.Completed = *body.Completed
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
		task := tasks.TaskResponse{
			Id:          &tsk.ID,
			Title:       &tsk.Title,
			Description: &tsk.Description,
			Completed:   &tsk.Completed,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) GetTasksId(_ context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	idInt, idIntErr := strconv.Atoi(request.Id)
	if idIntErr != nil {
		return nil, utils.EchoBadRequest(idIntErr)
	}

	task := new(Task)

	if err := h.service.GetById(task, uint(idInt)); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	response := tasks.GetTasksId200JSONResponse{
		Id:          &task.ID,
		Title:       &task.Title,
		Description: &task.Description,
		Completed:   &task.Completed,
	}

	return response, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	idInt, idIntErr := strconv.Atoi(request.Id)
	if idIntErr != nil {
		return nil, utils.EchoBadRequest(idIntErr)
	}

	task := new(Task)

	if err := h.service.GetById(task, uint(idInt)); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	body := request.Body

	if body.Title != nil {
		task.Title = *body.Title
	}

	if body.Description != nil {
		task.Description = *body.Description
	}

	if body.Completed != nil {
		task.Completed = *body.Completed
	}

	if err := h.service.Patch(task); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:          &task.ID,
		Title:       &task.Title,
		Description: &task.Description,
		Completed:   &task.Completed,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	idInt, idIntErr := strconv.Atoi(request.Id)
	if idIntErr != nil {
		return nil, utils.EchoBadRequest(idIntErr)
	}

	task := new(Task)

	if err := h.service.GetById(task, uint(idInt)); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	if err := h.service.Delete(task); err != nil {
		return nil, utils.EchoBadRequest(err)
	}

	response := tasks.DeleteTasksId200Response{}

	return response, nil
}
