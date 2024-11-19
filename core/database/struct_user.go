package database

type User struct {
	ID              string  `db:"id"`
	Name            string  `db:"name"`
	Email           string  `db:"email"`
	Role            string  `db:"role"`
	ProfileImageURL string  `db:"profile_image_url"`
	APIKey          *string `db:"api_key"`
	CreatedAt       int64   `db:"created_at"`
	UpdatedAt       int64   `db:"updated_at"`
	LastActiveAt    int64   `db:"last_active_at"`
	Settings        *string `db:"settings"`
	Info            *string `db:"info"`
	OAuthSub        *string `db:"oauth_sub"`
}
