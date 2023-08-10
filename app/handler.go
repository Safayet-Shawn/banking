package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/Safayet-Shawn/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode int64  `json:"zipcode"`
}
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("content-type") == "application/json" {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {

		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}
func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}
