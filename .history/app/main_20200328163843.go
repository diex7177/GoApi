package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "app/users"
	"strconv"

	"github.com/gorilla/mux"
)

var Users map[int]User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if _, ok := Users[user.Id]; ok {
		http.Error(w, "Already exists an user with the same id", http.StatusConflict)
		return
	} else {
		Users[user.Id] = user
	}
	log.Println("User: ", user, "added")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0)
	for _, value := range Users {
		userList = append(userList, value)
	}
	json.NewEncoder(w).Encode(userList)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user, ok := Users[id]; ok {
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user, ok := Users[id]; ok {
		delete(Users, id)
		log.Println("User: ", user, "removed")
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

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
