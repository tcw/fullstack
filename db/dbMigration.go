package db

import (
	"database/sql"
	"github.com/DavidHuie/gomigrate"
)

func MigrationUpdate(db *sql.DB, migrationPath string) {
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Sqlite3{}, migrationPath)
	if err != nil {
		panic("Migration setup failed")
	}
	err = migrator.Migrate()
	if err != nil {
		panic("Migration failed")
	}
}
