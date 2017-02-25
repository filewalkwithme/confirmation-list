package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func initializeWebServer() {
	http.HandleFunc("/confirmations", confirmationServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func confirmationServer(w http.ResponseWriter, req *http.Request) {
	if strings.ToUpper(req.Method) == "POST" {
		err := req.ParseForm()
		if err == nil {
			nameField, nameOk := req.PostForm["name"]
			confirmationCodeField, confirmationCodeOk := req.PostForm["confirmation-code"]
			companionsField, companionsOk := req.PostForm["companions"]
			if !(nameOk && confirmationCodeOk && companionsOk) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			name := nameField[0]
			confirmationCode := confirmationCodeField[0]
			companions := companionsField[0]

			fmt.Fprintf(w, "%v: %v: %v\n", name, confirmationCode, companions)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
