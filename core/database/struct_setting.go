package database

type Setting struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}
