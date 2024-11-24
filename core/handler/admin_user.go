package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func AdminUser(c echo.Context) error {
	req := &struct {
		Email string `form:"email" json:"email" query:"email"`
		Level string `form:"level" json:"level" `
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}

	switch c.Request().Method {
	case http.MethodPost:
		var users []database.User
		err := database.DB.Select(&users, "SELECT * FROM user")
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	case http.MethodPut:
		result, err := database.DB.Exec("UPDATE user SET level = ? WHERE email = ?", req.Level, req.Email)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	case http.MethodDelete:
		result, err := database.DB.Exec("DELETE FROM user WHERE email = ?", req.Email)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
