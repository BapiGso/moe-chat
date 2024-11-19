package database

type File struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Filename  string `db:"filename"`
	Meta      string `db:"meta"`
	CreatedAt int64  `db:"created_at"`
}
