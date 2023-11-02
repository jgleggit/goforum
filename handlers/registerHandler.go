package handlers

import (
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Print("Wrong Method. Use the POST method.")
		http.Error(w, "wrong method", http.StatusBadRequest)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "server couldn't parse the form",
		log.Printf("Couldn't Parse the Form"), StatusInternalServerError)
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("pasword")
	email := r.PostFormValue("email")
}