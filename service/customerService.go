package service

import (
	"github.com/pushrsp/banking/domain"
	"github.com/pushrsp/banking/dto"
	"github.com/pushrsp/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetTheCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return d.repo.FindAll(status)
}

func (d DefaultCustomerService) GetTheCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
