package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1000", "Brooke", "Los Angeles", "90028", "08-06-1995", "1"},
		{"1001", "Luca", "Santa Clarita", "91355", "01-10-2022", "1"},
	}
	return CustomerRepositoryStub{customers}
}
