package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	SaveUser(username string) sql.Result
}

type SqlRepository struct {
	db *sql.DB
}

func NewSqlRepository(sqlDb *sql.DB)  SqlRepository{
	return SqlRepository{sqlDb}
}

func (sr SqlRepository) SaveUser(username string) sql.Result {
	tx, err := sr.db.Begin()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := tx.Prepare("INSERT INTO userinfo(username) values(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	var res sql.Result
	for i := 0; i < 1000; i++ {
		res, err = stmt.Exec(username)
		if err != nil {
			panic("Could not execute insert")
		}
	}
	tx.Commit()

	return res;
}