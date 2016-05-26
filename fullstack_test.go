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
)

var (
	connection *sql.DB
	testport string
)

func TestMain(m *testing.M){
	connection = repository.NewMemoryDbConnection()
	db.MigrationUpdate(connection, "db/migrations")
	service := setupWebService(connection)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	testport = strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	go manners.Serve(listener, service)
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	user := repository.User{0, "my", "mylast"}
	buser, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:" + testport + "/add", "application/json", bytes.NewReader(buser))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fail()
	}
	t.Log(resp.StatusCode)
	manners.Close()
}