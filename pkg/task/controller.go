package task

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Controller(r *mux.Router, db *gorm.DB) {
	s := &Service{db}

	r.HandleFunc("/api/tasks", s.GetList).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks", s.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks/{id}", s.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks/{id}", s.Patch).Methods(http.MethodPatch)
	r.HandleFunc("/api/tasks/{id}", s.Delete).Methods(http.MethodDelete)
}
