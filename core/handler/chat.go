package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Chat(c echo.Context) error {
	req := &struct {
		UUID string `param:"UUID"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodGet:
		return c.Render(http.StatusOK, "handler.html", nil)
	}
	return echo.ErrMethodNotAllowed

}
