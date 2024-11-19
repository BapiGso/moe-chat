package database

type Knowledge struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Data        string `db:"data"` // Assuming JSON is stored as string
	Meta        string `db:"meta"` // Assuming JSON is stored as string
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
}
