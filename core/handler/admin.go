package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Admin(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return c.Render(http.StatusOK, "admin_header.html", nil)
	}
	return echo.ErrMethodNotAllowed

}
