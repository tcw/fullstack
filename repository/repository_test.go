package repository

import (
	"testing"
	"github.com/tcw/fullstack/db"
	"database/sql"
	"github.com/tcw/fullstack/domain"
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
	userRepo.SaveUser(domain.User{Username:"test1", Lastname:"test2"})
}

func TestGetUser(t *testing.T) {
	userRepo.SaveUser(domain.User{Username:"test", Lastname:"testesen"})
	user := userRepo.GetUser("test")
	if user.Lastname != "testesen" {
		t.Fatal("Couldn't get lastname")
	}
}