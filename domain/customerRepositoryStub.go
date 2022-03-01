package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1111", "Kang", "Seoul", "11110", "2021-01-01", "1"},
		{"2222", "Min", "Seoul", "11112", "2021-01-02", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
