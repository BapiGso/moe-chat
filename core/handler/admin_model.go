package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/api"
	"moechat/core/database"
	"net/http"
	"time"
)

func AdminModel(c echo.Context) error {
	req := &struct {
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
		// 将模型数组转换为对象
		modelMap := make(map[string]database.Model)
		for _, model := range models {
			modelMap[model.Type] = model
		}
		return c.JSON(http.StatusOK, modelMap)
	case http.MethodPut:
		req.CreatedAt = int(time.Now().Unix())
		_, err := database.DB.NamedExec(`INSERT OR REPLACE INTO model ( type, api_url, api_key, active, list, created_at)
VALUES ( :type, :api_url, :api_key, :active, :list, :created_at)`, database.Model(*req))
		if err != nil {
			return err
		}
		a := api.New(req.Type)
		return c.JSON(http.StatusOK, a.GetModelList())
	case http.MethodDelete:

	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
