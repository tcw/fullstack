# Fullstack - A 'fullstack' go project example

A minimal example of a Go rest service for the real world.

## Contains

*  Dependency manangement with vendoring (Godeps)
*  Sql database (Sqlite3)
*  Database migration (gomigrate)
*  Json rest service (Negroni, render)
*  Static http file server (Negroni)
*  TLS and Http2

Flags:

    -s, --skip-migration           Skip migration update
    -z, --profile                  Start cpu profiling
    -m, --memorydb                 Use in-memory database
    -f, --dbfile="./fullstack.db"  Database file
    -l, --reqlog                   Turn on request logging
    -p, --http="3000"              Web application port
    -t, --https="3001"             Secure web application port
    -k, --key="tls/key.pem"        Private Key
    -c, --cert="tls/cert.pem"      Public Key
        --version                  Show application version.

Post User to service

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d "{\"Username\":\"test\",\"Lastname\":\"tester\"}" http://localhost:3000/add

Get User from service

    curl localhost:3000/find/test


Used following command to create key and cert:

    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem