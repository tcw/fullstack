# Fullstack - A 'fullstack' go project example


## Contains

* Dependency manangement with vendor : Godeps
* Sql database: Sqlite3
* Database migration : gomigrate
* Json rest service : Negroni, render
* Static http file server : Negroni
* Graceful http server shutdown : manners

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d "{\"Username\":\"test\",\"Lastname\":\"tester\"}" http://localhost:3000/add
curl localhost:3000/find/test