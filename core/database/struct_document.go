package database

type Document struct {
	ID             int    `db:"id"`
	CollectionName string `db:"collection_name"`
	Name           string `db:"name"`
	Title          string `db:"title"`
	Filename       string `db:"filename"`
	Content        string `db:"content"`
	UserID         string `db:"user_id"`
	Timestamp      int    `db:"timestamp"`
}
