package repository

import (
	"testing"
	"github.com/tcw/go-graph/db"
	"database/sql"
)

var (
	userRepo UserDbRepository
	connection *sql.DB
)

func init() {
	connection = NewMemoryDbConnection()
	userRepo = NewUserRepository(connection)
	db.MigrationUpdate(connection, "../db/migrations")
}

func TestSaveUser(t *testing.T) {
	userRepo.SaveUser(User{Username:"test1", Lastname:"test2"})
}

func TestGetUser(t *testing.T) {
	userRepo.SaveUser(User{Username:"test", Lastname:"testesen"})
	user := userRepo.GetUser("test")
	if user.Lastname != "testesen" {
		t.Fatal("Couldn't get lastname")
	}
}