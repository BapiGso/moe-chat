package database

type Model struct {
	Name      string `db:"name"`
	Type      string `db:"type"`
	APIUrl    string `db:"api_url"`
	APIKey    string `db:"api_key"`
	Active    int    `db:"active"`
	List      string `db:"list"`
	CreatedAt int    `db:"created_at"`
}
