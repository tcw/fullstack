package repository

import (
	"testing"
	"github.com/tcw/go-graph/db"
)

func TestSaveUser(t *testing.T) {
	connection := NewMemoryDbConnection()
	repository := NewUserRepository(connection)
	db.MigrationUpdate(connection,"../db/migrations")
	repository.SaveUser("test")
}
