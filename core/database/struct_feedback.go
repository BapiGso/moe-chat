package database

type Feedback struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Version   int64  `db:"version"`
	Type      string `db:"type"`
	Data      string `db:"data"`     // Assuming JSON is stored as string
	Meta      string `db:"meta"`     // Assuming JSON is stored as string
	Snapshot  string `db:"snapshot"` // Assuming JSON is stored as string
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}
