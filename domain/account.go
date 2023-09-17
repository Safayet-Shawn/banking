package domain

import "github.com/Safayet-Shawn/banking/errs"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      string
	Status      string
}
type AccountRepository interface {
	Save(Account) (*Account, *errs.Apperror)
}
