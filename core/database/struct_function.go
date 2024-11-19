package database

type Function struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Name      string `db:"name"`
	Type      string `db:"type"`
	Content   string `db:"content"`
	Meta      string `db:"meta"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
	Valves    string `db:"valves"`
	IsActive  int    `db:"is_active"`
	IsGlobal  int    `db:"is_global"`
}
