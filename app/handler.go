package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Safayet-Shawn/banking/service"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode int64  `json:"zipcode"`
}
type CustomerHandler struct {
	service service.CustomerService
}

//	func greet(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Welcome to the Microservice Building with Golang !!")
//	}
func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "shawn", City: "dhaka", Zipcode: 1223},
	// 	{Name: "Suvo", City: "Sylet", Zipcode: 1209},
	// }
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("content-type") == "application/json" {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {

		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	v := mux.Vars(r)
// 	fmt.Println(v["customer_id"])
// 	fmt.Fprint(w, v["customer_id"])
// }
// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Created customer successfully")
// }
