package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toshi-pono/CORS-practice/server/handler"
)

func main() {
	handlers := handler.Handlers{}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"*"},
		AllowOriginFunc: func(origin string) (bool, error) {
			// どのOriginからでも受け付ける
			return true, nil
		},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{},
		AllowCredentials: true,
		ExposeHeaders:    []string{},
		MaxAge:           0,
	}))

	handlers.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
