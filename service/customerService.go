package service

import (
	"github.com/Safayet-Shawn/banking/domain"
	"github.com/Safayet-Shawn/banking/errs"
)

// primary port[all port are interfaces]
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.Apperror)
	GetCustomer(string) (*domain.Customer, *errs.Apperror)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.Apperror) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.Apperror) {
	return s.repo.ById(id)
}

// all new is helper function
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
