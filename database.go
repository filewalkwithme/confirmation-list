package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//DATABASEFILEPATH contains the path for the real sqlite3 db file
const DATABASEFILEPATH = "./confirmation-list.db"

func createDatabase() {
	if _, err := os.Stat(DATABASEFILEPATH); os.IsNotExist(err) {
		os.Remove(DATABASEFILEPATH)
		db, err := sql.Open("sqlite3", DATABASEFILEPATH)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		sqlStmt := `
create table guests (id integer not null primary key, name text, email text, confirmation_code text);
`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}
}
