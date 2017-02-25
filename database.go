package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//DATABASEFILEPATH contains the path for the real sqlite3 db file
const DATABASEFILEPATH = "./confirmation-list.db"

func initializeDB() (*sql.DB, error) {
	if _, err := os.Stat(DATABASEFILEPATH); os.IsNotExist(err) {
		return createDB()
	}

	db, err := sql.Open("sqlite3", DATABASEFILEPATH)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDB() (*sql.DB, error) {
	os.Remove(DATABASEFILEPATH)
	db, err := sql.Open("sqlite3", DATABASEFILEPATH)
	if err != nil {
		return nil, err
	}

	sqlStmt := `
create table guests (id integer not null primary key, name text, email text, confirmation_code text);
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
