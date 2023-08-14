package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
}

var users = []User{}

// GET /me
func (h *Handlers) GetMe(c echo.Context) error {
	return c.JSON(http.StatusOK, User{
		Username: c.Get("username").(string),
	})
}

// POST /users
func (h *Handlers) CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	users = append(users, user)
	return c.JSON(http.StatusCreated, user)
}

// POST /login
func (h *Handlers) Login(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	// passwordなしの適当なログイン機構
	for _, u := range users {
		if u.Username == user.Username {
			sess, err := session.Get("SESSION", c)
			if err != nil {
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			sess.Values["username"] = user.Username
			if err := sess.Save(c.Request(), c.Response()); err != nil {
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			return c.JSON(http.StatusOK, user)
		}
	}
	return echo.NewHTTPError(http.StatusUnauthorized)
}

// POST /logout
func (h *Handlers) Logout(c echo.Context) error {
	sess, err := session.Get("SESSION", c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	sess.Values["username"] = nil
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
