package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Users map[int]User

func main() {
	Users = make(map[int]User)
	log.Println("Default Users: ", Users)
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users{id}", GetUser).Methods("GET")
	router.HandleFunc("/users{id}", DeleteUser).Methods("DELETE")
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
