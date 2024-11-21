package database

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Title     string    `db:"title"`
	ShareID   string    `db:"share_id"`
	Archived  int       `db:"archived"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Messages  string    `db:"messages"` // Assuming JSON is stored as string
	Pinned    bool      `db:"pinned"`
	Meta      string    `db:"meta"` // Assuming JSON is stored as string
	FolderID  string    `db:"folder_id"`
}

type Message struct {
	Text    string `json:"text"`
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
