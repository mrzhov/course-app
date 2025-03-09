package modules

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/mrzhov/course-app/internal/modules/task"
	"github.com/mrzhov/course-app/internal/modules/user"
)

func Init(e *echo.Echo, db *gorm.DB) {
	g := e.Group("/api")

	task.Module(g, db)
	user.Module(g, db)
}
