package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
	"time"
)

// Completion 用于将会话发送给ai todo 可能需要池化
func Completion(c echo.Context) error {
	req := &struct {
		ID        uuid.UUID `param:"UUID"`
		Email     string
		Title     string
		ShareID   string
		Archived  int
		CreatedAt time.Time
		UpdatedAt time.Time
		Messages  string ` form:"message"` // Assuming JSON is stored as string
		Pinned    bool   ``
		Meta      string `` // Assuming JSON is stored as string
		FolderID  string ``
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Connection", "keep-alive")

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-c.Request().Context().Done():
				//会话结束把流式响应完的结果更新到数据库，req.uuid为空则生成uuid然后insert or replace
				if req.ID == uuid.Nil {
					req.ID = uuid.New()
				}
				_, err := database.DB.NamedExec(`INSERT OR REPLACE INTO 
  		  chat (id, email, title, share_id, archived, created_at, updated_at, messages, folder_id)
          VALUES (:id, :email, :title, :share_id, :archived, :created_at, :updated_at, :messages, :folder_id)`,
					database.Chat(*req))
				if err != nil {
					return err
				}
			case <-ticker.C:
				if _, err := fmt.Fprintf(w, "data: %s\n\n", []byte("time: "+time.Now().Format(time.RFC3339Nano))); err != nil {
					return err
				}
				w.Flush()
			}
		}
	case http.MethodPut:
		//case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
