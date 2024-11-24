package database

type File struct {
	Hash      string `db:"hash"`
	Email     string `db:"email"`
	Filename  string `db:"filename"`
	MimeType  string `db:"mime_type"`
	Data      []byte `db:"data"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}
