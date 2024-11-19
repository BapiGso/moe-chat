package database

type Folder struct {
	ID         string `db:"id"`
	ParentID   string `db:"parent_id"`
	UserID     string `db:"user_id"`
	Name       string `db:"name"`
	Items      string `db:"items"` // Assuming JSON is stored as string
	Meta       string `db:"meta"`  // Assuming JSON is stored as string
	IsExpanded bool   `db:"is_expanded"`
	CreatedAt  int64  `db:"created_at"`
	UpdatedAt  int64  `db:"updated_at"`
}
