package database

type Tool struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Name      string `db:"name"`
	Content   string `db:"content"`
	Specs     string `db:"specs"`
	Meta      string `db:"meta"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
	Valves    string `db:"valves"`
}
