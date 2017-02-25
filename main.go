package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type guest struct {
	name             string
	email            string
	confirmationCode string
}

func main() {
	db, err := initializeDB()
	if err != nil {
		panic("")
	}
	defer db.Close()

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
