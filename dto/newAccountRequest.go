package dto

import (
	"strings"

	"github.com/Safayet-Shawn/banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.Apperror {
	if r.Amount < 5000 {
		return errs.NewValidationError("Unable to open account because deposited amount less than 5000")
	} else if strings.ToLower(r.AccountType) != "saving" || strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type must be saving / checking")
	}
	return nil
}
