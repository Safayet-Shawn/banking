package domain

import (
	"github.com/Safayet-Shawn/banking/dto"
	"github.com/Safayet-Shawn/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}
type AccountRepository interface {
	Save(Account) (*Account, *errs.Apperror)
	FindBy(string) (*Account, *errs.Apperror)
	SaveTransaction(Transaction) (*Transaction, *errs.Apperror)
}

func (a Account) ToNewAccountResonseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}
func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
