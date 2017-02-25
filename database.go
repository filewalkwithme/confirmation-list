package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//DATABASEFILEPATH contains the path for the real sqlite3 db file
const DATABASEFILEPATH = "./confirmation-list.db"

var db *sql.DB

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
create table guests (id integer not null primary key, name text, email text, confirmation_code text, confirmation_date datetime, confirmed integer, companions integer);
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func insertGuest(db *sql.DB, guest guest) error {
	_, err := db.Exec("INSERT INTO guests (name, email, confirmation_code) VALUES (?, ?, ?)", guest.name, guest.email, guest.confirmationCode)
	if err != nil {
		return err
	}
	return nil
}

func confirmGuest(db *sql.DB, guest guest) error {
	count := 0
	err := db.QueryRow("select count(1) from guests where name=? and confirmation_code=? and confirmed=1", guest.name, guest.confirmationCode).Scan(&count)
	log.Printf("name: %v", guest.name)
	log.Printf("code: %v", guest.confirmationCode)
	log.Printf("count: %v", count)

	if count > 0 {
		return fmt.Errorf("Confirmation code already used.")
	}

	_, err = db.Exec("UPDATE guests set confirmed=1, confirmation_date=?, companions=? where name=? and confirmation_code=?", time.Now(), guest.companions, guest.name, guest.confirmationCode)
	if err != nil {
		return err
	}
	return nil
}

func loadGuestsFromCSV(db *sql.DB) {
	guestsFile, err := os.Open("guests.csv")
	if err != nil {
		log.Fatal(err)
	}

	var guests []guest

	guestsScanner := bufio.NewScanner(guestsFile)
	for guestsScanner.Scan() {
		line := guestsScanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) != 2 {
			log.Fatalf("Error reading line: %v", line)
		}
		guests = append(guests, guest{name: fields[0], email: fields[1], confirmationCode: generateConfirmationCode()})
	}

	for _, guest := range guests {
		log.Printf("Name: %v, Email: %v", guest.name, guest.email)
		err = insertGuest(db, guest)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := guestsScanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}
}
