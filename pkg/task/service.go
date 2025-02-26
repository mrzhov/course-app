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

func (s Service) Create(w http.ResponseWriter, r *http.Request) {
	body := &TaskBody{}
	bodyErr := _ComputeBody(w, r, body)

	if bodyErr != nil {
		return
	}

	var task Task
	task.Title = *body.Title

	if body.IsDone != nil {
		task.IsDone = *body.IsDone
	} else {
		task.IsDone = false
	}

	if res := s.db.Create(&task); res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, res.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (s Service) GetById(w http.ResponseWriter, r *http.Request) {
	task, err := s._GetById(w, r)

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*task)
}

func (s Service) Patch(w http.ResponseWriter, r *http.Request) {
	body := &TaskBody{}
	bodyErr := _ComputeBody(w, r, body)
	task, taskErr := s._GetById(w, r)

	if (bodyErr != nil) || (taskErr != nil) {
		return
	}

	task.Title = *body.Title

	if body.IsDone != nil {
		task.IsDone = *body.IsDone
	}

	s.db.Save(task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*task)
}

func (s Service) Delete(w http.ResponseWriter, r *http.Request) {
	task, err := s._GetById(w, r)

	if err != nil {
		return
	}

	s.db.Delete(task)
	fmt.Fprint(w, "Deleted successfully!")
}
