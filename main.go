package main

import (
	"flag"
)

type guest struct {
	name             string
	email            string
	confirmationCode string
	companions       int
}

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "", "load-from-csv, webserver")
	flag.Parse()

	dbLocal, err := initializeDB()
	if err != nil {
		panic("")
	}
	defer dbLocal.Close()
	db = dbLocal

	if mode == "load-from-csv" {
		loadGuestsFromCSV(db)
		return
	}

	if mode == "webserver" {
		initializeWebServer()
	}
}
