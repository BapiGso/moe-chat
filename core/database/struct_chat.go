package database

import (
	"time"
)

type Chat struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Title     string    `db:"title"`
	ShareID   string    `db:"share_id"`
	Archived  int       `db:"archived"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Chat      string    `db:"chat"` // Assuming JSON is stored as string
	Pinned    bool      `db:"pinned"`
	Meta      string    `db:"meta"` // Assuming JSON is stored as string
	FolderID  string    `db:"folder_id"`
}
