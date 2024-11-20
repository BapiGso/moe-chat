package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func AdminSetting(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
		var settings []database.Setting
		err := database.DB.Select(&settings, `SELECT * FROM setting`)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, settings)
	case http.MethodPut:
		exec, err := database.DB.Exec(`INSERT OR REPLACE INTO setting (key,value)
		VALUES (?,?)`, nil, nil)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, exec)
	case http.MethodDelete:
	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
