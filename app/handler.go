package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode int64  `json:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Microservice Building with Golang !!")
}
func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "shawn", City: "dhaka", Zipcode: 1223},
		{Name: "Suvo", City: "Sylet", Zipcode: 1209},
	}
	if r.Header.Get("content-type") == "application/json" {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {

		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Println(v["customer_id"])
	fmt.Fprint(w, v["customer_id"])
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Created customer successfully")
}
