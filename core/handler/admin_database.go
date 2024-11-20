package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AdminDatabase(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:

	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
