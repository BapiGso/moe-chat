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
func Chat(c echo.Context) error {
	req := &struct {
		ID        uuid.UUID `param:"id" json:"id"`
		Email     string
		Title     string `json:"Title"`
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
		if req.ID == uuid.Nil {
			return c.JSON(http.StatusOK, database.Chat{ID: uuid.New(), Messages: json.RawMessage("[]")})
		}
		var chat database.Chat
		err := database.DB.Get(&chat,
			"SELECT * FROM chat WHERE id = ? ", req.ID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return c.JSON(http.StatusOK, chat)
	case http.MethodPut: //用于更新会话内容
		req.Email = c.Get("email").(string)
		req.UpdatedAt = time.Now()
		_, err := database.DB.NamedExec(`INSERT OR REPLACE INTO
			 chat (id, email, title, share_id, archived, 
			       created_at, updated_at, messages, folder_id)
			VALUES (:id, :email, :title, :share_id, :archived, 
			        :created_at, :updated_at, :messages, :folder_id)`, database.Chat(*req))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, req)
	case http.MethodGet:
		return c.Render(http.StatusOK, "index.html", nil)
	}
	return echo.ErrMethodNotAllowed
}
