package task

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func (s Service) GetList(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	if res := s.db.Find(&tasks); res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, res.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

type CreateBody struct {
	Task struct {
		Title  string `json:"title"`
		IsDone bool   `json:"is_done"`
	} `json:"task"`
}

func (s Service) Create(w http.ResponseWriter, r *http.Request) {
	body := CreateBody{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if body.Task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Field \"task.title\" is required!")
		return
	}

	var task Task
	task.Title = body.Task.Title
	task.IsDone = body.Task.IsDone

	if res := s.db.Create(&task); res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, res.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
