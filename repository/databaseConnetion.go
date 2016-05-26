package repository

import "database/sql"

func NewDbConnection(file string) *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic("Could not open database connection")
	}
	return db
}

func NewMemoryDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Could not open database connection")
	}
	return db
}