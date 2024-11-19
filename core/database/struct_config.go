package database

import (
	"time"
)

type Config struct {
	ID        int       `db:"id"`
	Data      string    `db:"data"` // Assuming JSON is stored as string
	Version   int       `db:"version"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
}
