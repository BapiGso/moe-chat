package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"moechat/core/api/github"
	"moechat/core/database"
	"net/http"
	"time"
)

// Completion 用于将会话发送给ai todo 可能需要池化
// 前端要发的数据，什么公司的什么模型model，
// 什么对话message，uuid（如果没有则会响应完后把message插入数据库，有则会把之前的message带着这一次的replace数据库
func Completion(c echo.Context) error {
	req := &struct {
		Model string `form:"model"`
		Chat  struct {
			ID        uuid.UUID `param:"UUID"`
			Email     string
			Title     string
			ShareID   string
			Archived  int
			CreatedAt time.Time
			UpdatedAt time.Time
			Messages  string ` form:"message"`
			Pinned    bool   ``
			Meta      string ``
			FolderID  string ``
		}
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		reader, err := github.CreateOpenAIResponseReader(c.Request().Context())
		if err != nil {
			return err
		}
		buf := make([]byte, 512)
		for {
			select {
			case <-c.Request().Context().Done():
				//会话结束把流式响应完的结果更新到数据库，req.uuid为空则生成uuid然后insert or replace
				if req.Chat.ID == uuid.Nil {
					req.Chat.ID = uuid.New()
				}
				_, err := database.DB.NamedExec(`INSERT OR REPLACE INTO
			 chat (id, email, title, share_id, archived, created_at, updated_at, messages, folder_id)
			VALUES (:id, :email, :title, :share_id, :archived, :created_at, :updated_at, :messages, :folder_id)`, database.Chat(req.Chat))
				if err != nil {
					return err
				}
			default:
				n, err := reader.Read(buf)
				if err != nil {
					if err == io.EOF {
						return nil
					}
					return err
				}
				if n > 0 {
					fmt.Fprintf(w, "data: %s\n\n", buf[:n])
					w.Flush()
				}
			}

		}
	case http.MethodPut:
		//case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
