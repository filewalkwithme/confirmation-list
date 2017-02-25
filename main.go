package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	/*
		db, err := initializeDB()
		if err != nil {
			panic("")
		}
		defer db.Close()
	*/

	guestsFile, err := os.Open("guests.csv")
	if err != nil {
		log.Fatal(err)
	}

	guestsScanner := bufio.NewScanner(guestsFile)
	for guestsScanner.Scan() {
		log.Println(guestsScanner.Text()) // Println will add back the final '\n'
	}
	if err := guestsScanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}
}
