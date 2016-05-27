package main

import (
	"testing"
	"github.com/tcw/fullstack/repository"
	"database/sql"
	"github.com/tcw/fullstack/db"
	"github.com/braintree/manners"
	"net/http"
	"encoding/json"
	"bytes"
	"net"
	"os"
	"strconv"
	"github.com/tcw/fullstack/domain"
	"errors"
)

var (
	connection *sql.DB
	testport string
)

func TestMain(m *testing.M) {
	connection = repository.NewMemoryDbConnection()
	//connection = repository.NewDbConnection("./test.db")
	db.MigrationUpdate(connection, "db/migrations")
	service := setupRestService(connection)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	testport = strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	go manners.Serve(listener, service)
	code := m.Run()
	manners.Close()
	os.Exit(code)
}

func addUser(user domain.User) error {
	buser, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:" + testport + "/add", "application/json", bytes.NewReader(buser))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return errors.New("User not Created")
	}
	return nil
}

func TestAddingUserToRestService(t *testing.T) {
	user := domain.User{0, "my", "mylast"}
	err := addUser(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindingUserWithRestService(t *testing.T) {
	user := domain.User{0, "my2", "mylast2"}
	err := addUser(user)
	resp, err := http.Get("http://localhost:" + testport + "/find/my2")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}
	decoder := json.NewDecoder(resp.Body)
	var userRes []domain.User
	err = decoder.Decode(&userRes)
	if err != nil || userRes[0].Lastname != "mylast2" {
		t.Fail()
	}
}

func BenchmarkAddUser(b *testing.B) {
	user := domain.User{0, "my2", "mylast2"}
	buser, _ := json.Marshal(user)
	for i := 0; i < b.N; i++ {
		_, err := http.Post("http://localhost:" + testport + "/add", "application/json", bytes.NewReader(buser))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFindUser(b *testing.B) {
	user := domain.User{0, "my3", "mylast3"}
	err := addUser(user)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_ , err := http.Get("http://localhost:" + testport + "/find/my3")
		if err != nil {
			b.Fatal(err)
		}
	}
}
