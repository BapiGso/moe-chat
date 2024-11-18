package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return c.Render(http.StatusOK, "login.html", nil)
	}
	return echo.ErrMethodNotAllowed
}
