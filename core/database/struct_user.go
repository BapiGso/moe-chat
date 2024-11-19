package database

type User struct {
	Email           string `db:"email"`
	Password        []byte `db:"password"`
	Level           string `db:"level"`
	ProfileImageURL string `db:"profile_image_url"`
	CreatedAt       int    `db:"created_at"`
	UpdatedAt       int    `db:"updated_at"`
	Settings        string `db:"settings"`
}

const (
	LevelPending = "pending"
	LevelUser    = "user"
	LevelAdmin   = "admin"
)
