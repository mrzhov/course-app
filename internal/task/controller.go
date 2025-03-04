package task

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Controller(g *echo.Group, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gg := g.Group("/tasks")
	gg.GET("", handler.GetList)
	gg.POST("", handler.Create)
}

// r.HandleFunc("/api/tasks/{id}", controller.GetById).Methods(http.MethodGet)
// r.HandleFunc("/api/tasks/{id}", controller.Patch).Methods(http.MethodPatch)
// r.HandleFunc("/api/tasks/{id}", controller.Delete).Methods(http.MethodDelete)
