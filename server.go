package main

import (
	"log"
	"net/http"
	"strconv"
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

			companions, err := strconv.Atoi(companionsField[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			guest := guest{name: name, confirmationCode: confirmationCode, companions: companions}
			err = confirmGuest(db, guest)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
