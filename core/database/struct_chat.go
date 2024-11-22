package database

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID        uuid.UUID       `db:"id"`
	Email     string          `db:"email"`
	Title     string          `db:"title"`
	ShareID   string          `db:"share_id"`
	Archived  int             `db:"archived"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
	Messages  json.RawMessage `json:"messages"` // 使用 json.RawMessage 类型存储 JSON 字符串
	Pinned    bool            `db:"pinned"`
	Meta      string          `db:"meta"` // Assuming JSON is stored as string
	FolderID  string          `db:"folder_id"`
}
