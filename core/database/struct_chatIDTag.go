package database

type ChatIDTag struct {
	ID        string `db:"id"`
	TagName   string `db:"tag_name"`
	ChatID    string `db:"chat_id"`
	UserID    string `db:"user_id"`
	Timestamp int    `db:"timestamp"`
}
