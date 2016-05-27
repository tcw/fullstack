# Fullstack - A 'fullstack' go project example

A minimal example of a Go rest service for the real world.

## Contains

* Dependency manangement with vendoring : Godeps
* Sql database: Sqlite3
* Database migration : gomigrate
* Json rest service : Negroni, render
* Static http file server : Negroni
* Graceful http server shutdown : manners


    usage: fullstack [<flags>]

    Flags:
          --help                     Show context-sensitive help (also try
                                     --help-long and --help-man).
      -s, --skip-migration           Skip migration update
      -c, --profile                  Start cpu profiling
      -p, --port="3000"              Web application port
      -m, --memorydb                 Use in memory database
      -f, --dbfile="./fullstack.db"  Database file
      -l, --reqlog                   Turn on request logging
          --version                  Show application version.


curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d "{\"Username\":\"test\",\"Lastname\":\"tester\"}" http://localhost:3000/add
curl localhost:3000/find/test