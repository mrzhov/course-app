package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mrzhov/course-app/internal/task"
	"gorm.io/gorm"
)

func Register(e *echo.Echo, db *gorm.DB) {
	g := e.Group("/api")

	task.Module(g, db)
}
