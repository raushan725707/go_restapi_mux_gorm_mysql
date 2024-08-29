package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializerouter() {
	r := mux.NewRouter()
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/update", updateUser).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/get/{id}", getUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
