package domain

import "github.com/Safayet-Shawn/banking/errs"

// business object
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// secondery port
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.Apperror)
	ById(string) (*Customer, *errs.Apperror)
}
