package service

import (
	"github.com/Safayet-Shawn/banking/domain"
	"github.com/Safayet-Shawn/banking/errs"
)

// primary port[all port are interfaces]
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.Apperror)
	GetCustomer(string) (*domain.Customer, *errs.Apperror)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.Apperror) {
	return s.repo.FindAll()
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.Apperror) {
	return s.repo.ById(id)
}

// all new is helper function
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
