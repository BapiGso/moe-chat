package database

type Model struct {
	ID          string `db:"id"`
	Email       string `db:"email"`
	BaseModelID string `db:"base_model_id"`
	Name        string `db:"name"`
	Meta        string `db:"meta"`
	Params      string `db:"params"`
	CreatedAt   int    `db:"created_at"`
	UpdatedAt   int    `db:"updated_at"`
}
