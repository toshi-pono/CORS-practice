package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("SESSION", c)
			if err != nil {
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if sess.Values["username"] == nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}
			c.Set("username", sess.Values["username"].(string))
			return next(c)
		}
	}
}
