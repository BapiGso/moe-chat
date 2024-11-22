package database

type Model struct {
	Provider  string `db:"provider"`
	APIUrl    string `db:"api_url"`
	APIKey    string `db:"api_key"`
	Active    int    `db:"active"`
	List      string `db:"list"`
	CreatedAt int    `db:"created_at"`
}
