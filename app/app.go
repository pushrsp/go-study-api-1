package app

import (
	"github.com/gorilla/mux"
	"github.com/pushrsp/banking/domain"
	"github.com/pushrsp/banking/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")

	log.Fatalln(http.ListenAndServe(":8000", router))
}
