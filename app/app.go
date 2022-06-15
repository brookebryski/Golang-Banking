package app

import (
	"Golang-Banking/domain"
	"Golang-Banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// starting server
	log.Fatalln(http.ListenAndServe("localhost:8000", router))
}
