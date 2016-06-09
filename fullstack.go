package main

import (
	"database/sql"
	"fmt"
	"github.com/braintree/manners"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/tcw/fullstack/db"
	"github.com/tcw/fullstack/repository"
	"github.com/tcw/fullstack/web"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
)

var (
	skipMigration = kingpin.Flag("skip-migration", "Skip migration update").Short('s').Default("false").Bool()
	cpuProfile = kingpin.Flag("profile", "Start cpu profiling").Short('z').Default("false").Bool()
	memorydb = kingpin.Flag("memorydb", "Use in-memory database").Short('m').Default("false").Bool()
	dbfile = kingpin.Flag("dbfile", "Database file").Short('f').Default("./fullstack.db").String()
	requestlogging = kingpin.Flag("reqlog", "Turn on request logging").Short('l').Default("false").Bool()
	onlyhttp = kingpin.Flag("onlyhttp", "If true only http 1.1 will be used").Short('u').Default("false").Bool()
	siteRootHttps = kingpin.Flag("site-rooot", "Https site root").Short('r').Default("https://localhost:3001").String()
	httpPort = kingpin.Flag("http", "Http port").Short('p').Default("3000").String()
	httpsPort = kingpin.Flag("https", "Https port").Short('t').Default("3001").String()
	tlsKey = kingpin.Flag("key", "Tls Private Key (key)").Short('k').Default("tls/key.pem").String()
	tlsCert = kingpin.Flag("cert", "Tls Public Key (cert)").Short('c').Default("tls/cert.pem").String()
)

func main() {
	kingpin.Version(GetVersion())
	kingpin.Parse()

	//Database migrations
	var connection *sql.DB
	if *memorydb {
		connection = repository.NewMemoryDbConnection()
	} else {
		connection = repository.NewDbConnection(*dbfile)
	}

	if !*skipMigration {
		db.MigrationUpdate(connection, "./db/migrations")
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
	if *requestlogging {
		n.Use(negroni.NewLogger())
	}
	if *onlyhttp {
		log.Printf("Serving http redirect from port: %s\n", *httpPort)

			manners.ListenAndServe(":" + *httpPort, n)
	} else {
		log.Printf("Serving https from port: %s\n", *httpsPort)
		go func() {
			err := http.ListenAndServeTLS(":" + *httpsPort, *tlsCert, *tlsKey, n)
			if err != nil {
				log.Fatal("Https web server:", err)
			}
		}()
		log.Printf("Serving http redirect from port: %s\n", *httpPort)
		log.Fatal("Http redirct web server:",
			http.ListenAndServe(":" + *httpPort,
				http.RedirectHandler(*siteRootHttps, http.StatusMovedPermanently)))
	}

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
