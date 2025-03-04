package task

import (
	"github.com/labstack/echo/v4"
	"github.com/mrzhov/course-app/internal/web/tasks"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)
}

// r.HandleFunc("/api/tasks/{id}", controller.GetById).Methods(http.MethodGet)
// r.HandleFunc("/api/tasks/{id}", controller.Patch).Methods(http.MethodPatch)
// r.HandleFunc("/api/tasks/{id}", controller.Delete).Methods(http.MethodDelete)
