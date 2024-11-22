package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

// Chat get返回静态页面，post返回某个UUID对应的历史会话，put更新某个UUID的历史会话
func Chat(c echo.Context) error {
	req := &struct {
		UUID   string `param:"UUID"`
		Prompt string `form:"prompt"`
		Model  string `form:"model"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		if req.UUID == "" {
			return c.JSON(http.StatusOK, []struct{}{})
		}
		var chats []database.Chat
		if err := database.DB.Select(&chats, "SELECT * FROM chat WHERE id = ? ", req.UUID); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, chats)
	case http.MethodPut: //用于修改会话内容

	case http.MethodGet:
		return c.Render(http.StatusOK, "index.html", nil)
	}
	return echo.ErrMethodNotAllowed
}
