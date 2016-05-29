package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tcw/fullstack/domain"
	"log"
)

type UserRepository interface {
	SaveUser(user domain.User) sql.Result
	GetUser(username string) []domain.User
}

type UserDbRepository struct {
	db *sql.DB
}

func NewUserRepository(sqlDb *sql.DB) UserDbRepository {
	return UserDbRepository{sqlDb}
}

func (sr UserDbRepository) SaveUser(user domain.User) sql.Result {
	tx, err := sr.db.Begin()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := tx.Prepare("INSERT INTO userinfo(firstname,lastname) values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(user.Firstname, user.Lastname)
	if err != nil {
		panic("Could not execute insert")
	}
	tx.Commit()
	return res
}

func (sr UserDbRepository) GetUser(firstname string) []domain.User {
	var uid int64
	var fname string
	var lastname string
	rows, err := sr.db.Query("select uid,firstname,lastname from userinfo where firstname = ?", firstname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		err := rows.Scan(&uid, &fname, &lastname)
		users = append(users, domain.User{uid, fname, lastname})
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users
}
