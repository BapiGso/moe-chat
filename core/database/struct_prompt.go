package database

type Prompt struct {
	ID        int    `db:"id"`
	Command   string `db:"command"`
	UserID    string `db:"user_id"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	Timestamp int    `db:"timestamp"`
}
