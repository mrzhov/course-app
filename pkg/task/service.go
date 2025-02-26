package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

func (s Service) Create(w http.ResponseWriter, r *http.Request) {
	body := CreateBody{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if body.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Field \"title\" is required!")
		return
	}

	var task Task
	task.Title = body.Title
	task.IsDone = body.IsDone

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
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	var task Task

	if res := s.db.First(&task, id); res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, res.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

type PatchBody struct {
	Title  *string `json:"title"`
	IsDone *bool   `json:"is_done"`
}

func (s Service) Patch(w http.ResponseWriter, r *http.Request) {
	id, idErr := strconv.Atoi(mux.Vars(r)["id"])
	body := PatchBody{}

	if idErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, idErr)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if body.Title == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Field \"title\" is required!")
		return
	}

	var task Task

	if res := s.db.First(&task, id); res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, res.Error)
		return
	}

	task.Title = *body.Title

	if body.IsDone != nil {
		task.IsDone = *body.IsDone
	}

	s.db.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (s Service) Delete(w http.ResponseWriter, r *http.Request) {
	id, idErr := strconv.Atoi(mux.Vars(r)["id"])

	if idErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, idErr)
		return
	}

	var task Task

	if res := s.db.First(&task, id); res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, res.Error)
		return
	}

	s.db.Delete(&task)

	fmt.Fprint(w, "Deleted successfully!")
}
