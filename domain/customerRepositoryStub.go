package domain

// adapted [mock adapter]
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// all new is helper function
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Shawn", City: "Dhaka", Zipcode: "1203", DateOfBirth: "21-02-1996"},
		{Id: "2", Name: "Suvo", City: "Narsingdi", Zipcode: "1104", DateOfBirth: "24-02-1995"},
	}
	return CustomerRepositoryStub{customers: customers}
}
