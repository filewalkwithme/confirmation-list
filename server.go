package main

import (
	"io"
	"log"
	"net/http"
)

func initializeWebServer() {
	http.HandleFunc("/confirmations", confirmationServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func confirmationServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Confirmed!\n")
}
