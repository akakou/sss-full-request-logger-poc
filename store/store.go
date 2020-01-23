package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/http"
	"os"
)

var dbfile string = "data/db.sqlite3"
var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Print(string(body))

	stmt, _ := db.Prepare(`INSERT INTO "payload" ("data") VALUES (?) `)
	stmt.Exec(string(body))
	stmt.Close()

}

func main() {
	os.Remove(dbfile)
	db, _ = sql.Open("sqlite3", dbfile)
	db.Exec(`CREATE TABLE "payload" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "data" VARCHAR(255))`)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
