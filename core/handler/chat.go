package handler

import (
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
)

func Chat(c echo.Context) error {
	req := &struct {
		UUID string `query:"UUID"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		var chats []database.Chat
		if err := database.DB.Select(&chats, "SELECT * FROM chat WHERE id = ? AND user_id", req.UUID, c.Get("UID")); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, chats)
	case http.MethodGet:
		return c.Render(http.StatusOK, "index.html", nil)
	}
	return echo.ErrMethodNotAllowed

}
