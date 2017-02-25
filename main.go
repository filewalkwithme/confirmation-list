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
	var mode string
	flag.StringVar(&mode, "mode", "", "load-from-csv, webserver")
	flag.Parse()

	db, err := initializeDB()
	if err != nil {
		panic("")
	}
	defer db.Close()

	if mode == "load-from-csv" {
		loadGuestsFromCSV(db)
		return
	}

	if mode == "webserver" {
		initializeWebServer()
	}
}
