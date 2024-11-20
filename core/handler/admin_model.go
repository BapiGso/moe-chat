package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
	"time"
)

func AdminModel(c echo.Context) error {
	req := &struct {
		Name      string `form:"name" json:"name" db:"id"`
		Type      string `form:"type" json:"type" db:"type"`
		APIUrl    string `form:"apiurl" json:"apiurl" db:"api_url"`
		APIKey    string `form:"apikey" json:"apikey" db:"api_key"`
		Active    int    `form:"active" json:"active" db:"active"`
		List      string `form:"list" json:"list" db:"list"`
		CreatedAt int    `db:"created_at"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		var models []database.Model
		err := database.DB.Select(&models, "SELECT * FROM model")
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models)
	case http.MethodPut:
		req.CreatedAt = int(time.Now().Unix())
		exec, err := database.DB.NamedExec(`INSERT OR REPLACE INTO model (name, type, api_url, api_key, active, list, created_at)
VALUES (:name, :type, :api_url, :api_key, :active, :list, :created_at)`, database.Model(*req))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, exec)
	case http.MethodDelete:

	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
