package main

import (
	"flag"
)

type guest struct {
	name             string
	email            string
	confirmationCode string
}

func main() {
	var populateDB bool
	flag.BoolVar(&populateDB, "populate-db", false, "")
	flag.Parse()

	db, err := initializeDB()
	if err != nil {
		panic("")
	}
	defer db.Close()

	if populateDB {
		loadGuestsFromCSV(db)
	}
}
