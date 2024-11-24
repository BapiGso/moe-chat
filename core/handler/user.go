package handler

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func User(c echo.Context) error {
	req := &struct {
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		var user database.User
		err := database.DB.Get(&user,
			"SELECT * FROM main.user WHERE email = ? ", c.Get("email"))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return c.JSON(http.StatusOK, user)
	case http.MethodPut: //用于更新会话内容

	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
