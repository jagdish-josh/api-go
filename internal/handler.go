package internal

import (
	"net/http"
	"encoding/json"
	"log"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	result := HealthService()
	w.Write([]byte(result))
}

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("home page triggered")
	w.Write([]byte("Welcome to the Homepage"))
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	result, err := GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}


func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	result:= CreateUser(r)
	log.Println("CreateUserHandler result:", result)
	w.Write([]byte(result))
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	result:= UpdateUser(r)
	w.Write([]byte(result))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	result:= DeleteUser(r)
	w.Write([]byte(result))
}
