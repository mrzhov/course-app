package task

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/mrzhov/course-app/internal/web/tasks"
)

func Module(g *echo.Group, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(g, strictHandler)

	// g := group.Group("/tasks")
	// g.POST("", handler.Create)
	// g.GET("", handler.GetList)
	// g.GET("/:id", handler.GetById)
	// g.PATCH("/:id", handler.Patch)
	// g.DELETE("/:id", handler.Delete)
}
