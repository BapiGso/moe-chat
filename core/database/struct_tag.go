package database

type Tag struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	UserID string `db:"user_id"`
	Meta   string `db:"meta"` // Assuming JSON is stored as string
}
