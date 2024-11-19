package database

import "time"

type MigrateHistory struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	MigratedAt time.Time `db:"migrated_at"`
}
