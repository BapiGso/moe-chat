package database

type Memory struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Content   string `db:"content"`
	UpdatedAt int    `db:"updated_at"`
	CreatedAt int    `db:"created_at"`
}
