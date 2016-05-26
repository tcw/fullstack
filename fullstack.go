package main

import (
	"os"
	"log"
	"fmt"
	"runtime/pprof"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tcw/fullstack/repository"
	"github.com/tcw/fullstack/db"
	"github.com/tcw/fullstack/web"
	"database/sql"
	"github.com/braintree/manners"
)

var (
	migration = kingpin.Flag("migrate", "Migraion command (update)").Short('m').Default("").String()
	cpuProfile = kingpin.Flag("profile", "Starts profiling").Short('c').Default("false").Bool()
	port = kingpin.Flag("port", "web application port").Short('p').Default("3000").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	//Database migrations
	connection := repository.NewDbConnection()
	if *migration == "update" {
		db.MigrationUpdate(connection,"./db/migrations")
	}

	//Cpu profiling
	//cmd:'go tool pprof go-graph cpu.pprof'
	if *cpuProfile {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		fmt.Println("Run 'go tool pprof go-graph cpu.pprof' to watch profile")
		fmt.Println("Profile is being written to 'cpu.pprof' ...")
		defer pprof.StopCPUProfile()
	}

	//Web server
	n := setupRestService(connection)
	manners.ListenAndServe(":"+ *port, n)
}

func setupRestService(conn *sql.DB) *negroni.Negroni {
	router := mux.NewRouter()
	sqlRepository := repository.NewUserRepository(conn)
	userRepository := web.NewUserWeb(sqlRepository)

	router.Handle("/add", userRepository.AddUserHandler()).Methods("POST")
	router.Handle("/find/{username}", userRepository.GetUserHandler()).Methods("GET")

	n := negroni.New(
		negroni.NewStatic(http.Dir("web/static")),
		negroni.NewRecovery())
	n.UseHandler(router)
	return n
}

