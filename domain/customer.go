package domain

import "github.com/Safayet-Shawn/banking/errs"

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
	FindAll(status string) ([]Customer, *errs.Apperror)
	ById(string) (*Customer, *errs.Apperror)
}
