package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//define route
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Println(v["customer_id"])
	fmt.Fprint(w, v["customer_id"])
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Created customer successfully")
}
