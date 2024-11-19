package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func User(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
		var users []database.User
		err := database.DB.Select(&users, "SELECT * FROM user")
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	case http.MethodGet:
		return c.Render(http.StatusOK, "handler.html", nil)
	}
	return echo.ErrMethodNotAllowed
}
