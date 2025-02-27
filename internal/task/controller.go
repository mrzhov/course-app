package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	service *Service
}

func NewContoller(service *Service) *Controller {
	return &Controller{service}
}

type CreateBody struct {
	Title  *string `json:"title"`
	IsDone bool    `json:"is_done"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Title == nil {
		err := errors.New("error: field \"title\" is required")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task Task
	task.Title = *body.Title
	task.IsDone = body.IsDone
	err := c.service.Create(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (c *Controller) GetList(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.service.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (c *Controller) GetById(w http.ResponseWriter, r *http.Request) {
	id, idErr := strconv.Atoi(mux.Vars(r)["id"])

	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusBadRequest)
		return
	}

	task, taskErr := c.service.GetById(uint(id))
	if taskErr != nil {
		http.Error(w, taskErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

type PatchBody struct {
	Title  *string `json:"title"`
	IsDone *bool   `json:"is_done"`
}

func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	id, idErr := strconv.Atoi(mux.Vars(r)["id"])
	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusBadRequest)
		return
	}

	var body PatchBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Title == nil && body.IsDone == nil {
		err := errors.New("error: at least one field must be specified")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, taskErr := c.service.GetById(uint(id))
	if taskErr != nil {
		http.Error(w, taskErr.Error(), http.StatusInternalServerError)
		return
	}

	if body.Title != nil {
		task.Title = *body.Title
	}
	if body.IsDone != nil {
		task.IsDone = *body.IsDone
	}

	patchErr := c.service.Patch(&task)

	if patchErr != nil {
		http.Error(w, patchErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, idErr := strconv.Atoi(mux.Vars(r)["id"])

	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusBadRequest)
		return
	}

	task, taskErr := c.service.GetById(uint(id))
	if taskErr != nil {
		http.Error(w, taskErr.Error(), http.StatusInternalServerError)
		return
	}

	deleteErr := c.service.Delete(&task)

	if deleteErr != nil {
		http.Error(w, deleteErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseMessage{Message: "Deleted successfully!"})
}
