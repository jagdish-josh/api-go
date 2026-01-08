package main

import (
	"log"
	"net/http"

	"api-go/internal"
)

func main() {
	router := internal.NewRouter()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
