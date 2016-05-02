package db

import (
	"github.com/DavidHuie/gomigrate"
	"database/sql"
)

func MigrationUpdate(db *sql.DB) {
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Sqlite3{}, "./db/migrations")
	if err != nil {
		panic("Migration setup failed")
	}
	err = migrator.Migrate()
	if err != nil {
		panic("Migration failed")
	}
}