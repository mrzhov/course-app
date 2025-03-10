package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Module(group *echo.Group, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	g := group.Group("/users")
	g.POST("", handler.Create)
	g.GET("", handler.GetList)

	gg := group.Group("/users/:id")
	gg.GET("", handler.GetById)
	gg.PATCH("", handler.Patch)
	gg.DELETE("", handler.Delete)
	gg.GET("/tasks", handler.GetTasks)
}
