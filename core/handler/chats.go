package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
	"time"
)

// Chat get返回静态页面，post返回某个UUID对应的历史会话，put更新某个UUID的历史会话
func Chats(c echo.Context) error {
	req := &struct {
		ID        uuid.UUID `param:"id" json:"id"`
		Email     string
		Title     string
		ShareID   string
		Archived  int
		CreatedAt time.Time
		UpdatedAt time.Time
		Messages  json.RawMessage `json:"messages"` // 使用 json.RawMessage
		Pinned    bool
		Meta      string
		FolderID  string
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		var chats []database.Chat
		err := database.DB.Select(&chats,
			"SELECT * FROM chat WHERE email = '' ", req.ID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return c.JSON(http.StatusOK, chats)
	case http.MethodPut: //用于修改会话内容
	case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
