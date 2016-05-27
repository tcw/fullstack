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
	migrationUpdate = kingpin.Flag("migrate", "Migraion command").Short('u').Default("true").Bool()
	cpuProfile = kingpin.Flag("profile", "Starts profiling").Short('c').Default("false").Bool()
	port = kingpin.Flag("port", "web application port").Short('p').Default("3000").String()
	memorydb = kingpin.Flag("memorydb", "web application port").Short('m').Default("false").Bool()
	dbfile = kingpin.Flag("dbfile", "web application port").Short('f').Default("./fullstack.db").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	//Database migrations
	var connection *sql.DB
	if *memorydb {
		connection= repository.NewMemoryDbConnection()
	}else {
		connection= repository.NewDbConnection(*dbfile)
	}

	if *migrationUpdate {
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
		fmt.Println("Run 'go tool pprof fullstack cpu.pprof' to watch profile")
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

