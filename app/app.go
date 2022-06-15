package app

import (
	"Golang-Banking"
	"log"
	"net/http"
)

func Start() {
	// define routes
	http.HandleFunc("/greet", main.greet)
	http.HandleFunc("/customers", main.getAllCustomers)

	// starting server
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}
