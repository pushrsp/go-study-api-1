package domain

import "github.com/pushrsp/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id string) (*Customer, *errs.AppError)
}
