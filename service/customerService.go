package service

import "github.com/Safayet-Shawn/banking/domain"

// primary port[all port are interfaces]
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// all new is helper function
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
