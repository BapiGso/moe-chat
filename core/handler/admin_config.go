package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func AdminConfig(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
		var config database.Config
		err := database.DB.Get(&config, `SELECT * FROM config WHERE key = ?`, c.QueryParam("key"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, config)
	case http.MethodPut:
		exec, err := database.DB.Exec(`INSERT OR REPLACE INTO config (key,val)
		VALUES (?,?)`, c.QueryParam("key"), c.QueryParam("val"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, exec)
	case http.MethodDelete:
	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
