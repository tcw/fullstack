package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type UserRepository interface {
	SaveUser(user User) sql.Result
	GetUser(username string) User
}

type UserDbRepository struct {
	db *sql.DB
}

func NewUserRepository(sqlDb *sql.DB) UserDbRepository {
	return UserDbRepository{sqlDb}
}

type User struct {
	Uid      int64
	Username string
	Lastname string
}

func (sr UserDbRepository) SaveUser(user User) sql.Result {
	tx, err := sr.db.Begin()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := tx.Prepare("INSERT INTO userinfo(username,lastname) values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(user.Username,user.Lastname)
	if err != nil {
		panic("Could not execute insert")
	}
	tx.Commit()
	return res;
}

func (sr UserDbRepository) GetUser(username string) User {
	var uid int64
	var uname string
	var lastname string
	rows,err := sr.db.Query("select uid,username,lastname from userinfo where username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&uid,&uname,&lastname)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return User{uid,uname,lastname};
}