package app

import (
	"Golang-Banking/domain"
	"Golang-Banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//func sanityCheck() {
//	if os.Getenv("SERVER_ADDRESS") == "" ||
//		os.Getenv("SERVER_PORT") == "" {
//		log.Fatalln("Environment variable not defined...")
//	}
//}

func Start() {

	//sanityCheck()

	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// starting server
	//address := os.Getenv("SERVER_ADDRESS")
	//port := os.Getenv("SERVER_PORT")
	//log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
	log.Fatalln(http.ListenAndServe("localhost:8000", router))
}
