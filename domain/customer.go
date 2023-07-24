package domain

// business object
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// secondery port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
