package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func AdminSetting(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
		var config []database.Config
		err := database.DB.Select(&config, `SELECT * FROM config`)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, config)
	case http.MethodPut:
		exec, err := database.DB.Exec(`INSERT OR REPLACE INTO config (key,value)
		VALUES (?,?)`, c.QueryParam("key"), c.QueryParam("value"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, exec)
	case http.MethodDelete:
	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
