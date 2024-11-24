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
		Provider  string `form:"provider" json:"provider" db:"provider"`
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
			modelMap[model.Provider] = model
		}
		return c.JSON(http.StatusOK, modelMap)
	case http.MethodPut:
		req.CreatedAt = int(time.Now().Unix())
		result, err := database.DB.NamedExec(`INSERT OR REPLACE INTO 
    		model (provider, api_url, api_key, active, list, created_at)
			VALUES ( :provider, :api_url, :api_key, :active, :list, :created_at)`, database.Model(*req))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	case http.MethodDelete:
	case http.MethodGet:
	case http.MethodOptions:
		a := api.New(req.Provider)
		modelList := a.GetModelList()
		return c.JSON(http.StatusOK, modelList)
	}
	return echo.ErrMethodNotAllowed
}
