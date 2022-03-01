package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pushrsp/banking/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zpi_code" xml:"zipcode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func (c *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := c.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (c *CustomerHandler) getTheCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := c.service.GetTheCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post received")
}
