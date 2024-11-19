package database

type File struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Filename  string `db:"filename"`
	Meta      string `db:"meta"` // Assuming JSON is stored as string
	CreatedAt int    `db:"created_at"`
	Hash      string `db:"hash"`
	Data      string `db:"data"` // Assuming JSON is stored as string
	UpdatedAt int64  `db:"updated_at"`
	Path      string `db:"path"`
}
