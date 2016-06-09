# Fullstack - A 'fullstack' go project example

A minimal example of a Go rest service for the real world.

## Contains

*  Dependency manangement with vendoring (Godeps)
*  Sql database (Sqlite3)
*  Database migration (gomigrate)
*  Json rest service (Negroni, render)
*  Static http file server (Negroni)
*  TLS and Http2 with http redirect

## Install

    go get -u github.com/tcw/fullstack

## Usage

Run example project:

    go get -u github.com/tools/godep
    cd $GOPATH/src/github.com/tcw/fullstack
    godep go build
    ./fullstack

Flags:

        --help                     Show context-sensitive help (also try
                             --help-long and --help-man).
    -s, --skip-migration           Skip migration update
    -z, --profile                  Start cpu profiling
    -m, --memorydb                 Use in-memory database
    -f, --dbfile="./fullstack.db"  Database file
    -l, --reqlog                   Turn on request logging
    -u, --onlyhttp                 If true only http 1.1 will be used
    -r, --site-rooot="https://localhost:3001"
                             Https site root
    -p, --http="3000"              Http port
    -t, --https="3001"             Https port
    -k, --key="tls/key.pem"        Tls Private Key (key)
    -c, --cert="tls/cert.pem"      Tls Public Key (cert)
        --version                  Show application version.

Post User to service (start with 'fullstack -u')

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d "{\"Firstname\":\"test\",\"Lastname\":\"tester\"}" http://localhost:3000/add

Get User from service (start with 'fullstack -u')

    curl localhost:3000/find/test

Key and cert in the tls folder was generated using:

    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem