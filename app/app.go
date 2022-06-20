package app

import (
	"Golang-Banking/domain"
	"Golang-Banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
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
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	//accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// starting server
	//address := os.Getenv("SERVER_ADDRESS")
	//port := os.Getenv("SERVER_PORT")
	//log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
	log.Fatalln(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	//dbUser := os.Getenv("DB_USER")
	//dbPasswd := os.Getenv("DB_PASSWD")
	//dbAddr := os.Getenv("DB_ADDR")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")
	//
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/$s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	//client, err := sqlx.Open("mysql", dataSource)
	client, err := sqlx.Open("mysql", "root:Bb080695@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
