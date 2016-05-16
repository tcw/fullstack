package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type UserRepository interface {
	SaveUser(username string) sql.Result
}

type UserDbRepository struct {
	db *sql.DB
}

func NewUserRepository(sqlDb *sql.DB) UserDbRepository {
	return UserDbRepository{sqlDb}
}

func (sr UserDbRepository) SaveUser(username string) sql.Result {
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
	res, err = stmt.Exec(username)
	if err != nil {
		panic("Could not execute insert")
	}
	tx.Commit()
	return res;
}