package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi-pono/CORS-practice/server/middleware"
)

type Handlers struct{}

func (h *Handlers) Register(e *echo.Echo) {
	tasks := e.Group("/tasks", middleware.AuthMiddleware())
	{
		tasks.GET("", h.GetTasks)
		tasks.POST("", h.CreateTask)
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}
	e.GET("/me", h.GetMe, middleware.AuthMiddleware())

	e.POST("/users", h.CreateUser)
	e.POST("/login", h.Login)
	e.POST("/logout", h.Logout)
}
