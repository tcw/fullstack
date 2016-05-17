package main

import (
	"github.com/tcw/go-graph/repository"
	"os"
	"log"
	"fmt"
	"runtime/pprof"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/tcw/go-graph/db"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/tcw/go-graph/web"
	"github.com/gorilla/mux"
)

var (
	migration = kingpin.Flag("migrate", "Migraion command (update)").Short('m').Default("").String()
	profile = kingpin.Flag("profile", "Starts profiling").Short('p').Default("false").Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	if *profile {
		startProfiling()
	}
	connection := repository.NewDbConnection()
	if *migration == "update" {
		db.MigrationUpdate(connection,"./db/migrations")
	}

	mux := mux.NewRouter()
	sqlRepository := repository.NewUserRepository(connection)
	userRepository := web.NewUserWeb(sqlRepository)

	mux.Handle("/add/{username}/{lastname}", userRepository.AddUserHandler())
	mux.Handle("/find/{username}", userRepository.GetUserHandler())

	n := negroni.New(negroni.NewStatic(http.Dir("web/static")),
		negroni.NewRecovery())
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}

func startProfiling() {
	//go tool pprof go-graph cpu.pprof
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	fmt.Println("Run 'go tool pprof go-graph cpu.pprof' to watch profile")
	fmt.Println("Profile is being written to 'cpu.pprof' ...")
	defer pprof.StopCPUProfile()
}
