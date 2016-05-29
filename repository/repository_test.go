package repository

import (
	"database/sql"
	"github.com/tcw/fullstack/db"
	"github.com/tcw/fullstack/domain"
	"testing"
)

var (
	userRepo   UserDbRepository
	connection *sql.DB
)

func init() {
	connection = NewMemoryDbConnection()
	userRepo = NewUserRepository(connection)
	db.MigrationUpdate(connection, "../db/migrations")
}

func TestSaveUser(t *testing.T) {
	userRepo.SaveUser(domain.User{Username: "test1", Lastname: "test2"})
}

func TestGetUser(t *testing.T) {
	userRepo.SaveUser(domain.User{Username: "test", Lastname: "testesen"})
	userRepo.SaveUser(domain.User{Username: "test", Lastname: "testesen3"})
	user := userRepo.GetUser("test")
	if len(user) != 2 {
		t.Fail()
	}
	if user[0].Lastname != "testesen" {
		t.Fatal("Couldn't get lastname")
	}
}
