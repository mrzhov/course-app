package task

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	service *Service
}

func NewContoller(service *Service) *Controller {
	return &Controller{service}
}

// func (s Service) Create(w http.ResponseWriter, r *http.Request) {
// 	body := &TaskBody{}
// 	bodyErr := _ComputeBody(w, r, body)

// 	if bodyErr != nil {
// 		return
// 	}

// 	var task Task
// 	task.Title = *body.Title

// 	if body.IsDone != nil {
// 		task.IsDone = *body.IsDone
// 	} else {
// 		task.IsDone = false
// 	}

// 	if res := s.db.Create(&task); res.Error != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprint(w, res.Error)
// 		return
// 	}
// }

// func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	var task taskService.Task
// 	err := json.NewDecoder(r.Body).Decode(&task)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	createdTask, err := h.Service.CreateTask(task)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(createdTask)
// }

type CreateBody struct {
	Title  *string `json:"title"`
	IsDone *bool   `json:"is_done"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(task)
	json.NewEncoder(w).Encode(body)
}

func (c *Controller) GetList(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetById(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {

}
