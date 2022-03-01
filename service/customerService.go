package service

import "github.com/pushrsp/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetTheCustomer(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return d.repo.FindAll()
}

func (d DefaultCustomerService) GetTheCustomer(id string) (*domain.Customer, error) {
	return d.repo.FindById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
