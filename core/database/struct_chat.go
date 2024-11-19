package database

import "time"

type Chat struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Title     string    `db:"title"`
	Chat      string    `db:"chat"`
	ShareID   *string   `db:"share_id"`
	Archived  bool      `db:"archived"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
