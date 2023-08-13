package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Safayet-Shawn/banking/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, _ := ch.service.GetAllCustomer(status)
	if r.Header.Get("content-type") == "application/json" {
		WriteResponse(w, http.StatusOK, customers)
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
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}
}
func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
