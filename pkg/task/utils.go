package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s Service) _GetById(w http.ResponseWriter, r *http.Request) (*Task, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return nil, err
	}

	var task Task

	if res := s.db.First(&task, id); res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "error: task not found")
		return nil, res.Error
	}

	return &task, nil
}

type TaskBody struct {
	Title  *string `json:"title"`
	IsDone *bool   `json:"is_done"`
}

func _ComputeBody(w http.ResponseWriter, r *http.Request, body *TaskBody) error {
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return err
	}

	if body.Title == nil {
		w.WriteHeader(http.StatusBadRequest)
		err := errors.New("error: field \"title\" is required")
		fmt.Fprint(w, err)
		return err
	}

	return nil
}
