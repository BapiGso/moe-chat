package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func Model(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
		var models []database.Model
		err := database.DB.Select(&models, `SELECT provider,list,active FROM model WHERE active = 1`)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models)
	}
	return echo.ErrMethodNotAllowed

}
