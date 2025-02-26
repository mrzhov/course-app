package task

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	controller := NewContoller(service)

	r.HandleFunc("/api/tasks", controller.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks", controller.GetList).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks/{id}", controller.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks/{id}", controller.Patch).Methods(http.MethodPatch)
	r.HandleFunc("/api/tasks/{id}", controller.Delete).Methods(http.MethodDelete)
}
