package app

import (
	"log"
	"net/http"

	"github.com/Safayet-Shawn/banking/domain"
	"github.com/Safayet-Shawn/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	//define route
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//wiring the application
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
