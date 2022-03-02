package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pushrsp/banking/service"
	"net/http"
)

//type Customer struct {
//	Name    string `json:"full_name" xml:"name"`
//	City    string `json:"city" xml:"city"`
//	Zipcode string `json:"zpi_code" xml:"zipcode"`
//}

type CustomerHandler struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func (c *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.service.GetAllCustomers(r.URL.Query().Get("status"))
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customers)
	}
}

func (c *CustomerHandler) getTheCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := c.service.GetTheCustomer(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post received")
}
