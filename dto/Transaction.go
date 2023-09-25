package dto

import "github.com/Safayet-Shawn/banking/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	Account_id      string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}
type TransactionResponse struct {
	TransactionId string `json:"transaction_id"`
	AccountId     string `json:"account_id"`
}

func (t TransactionRequest) Validate() *errs.Apperror {
	if t.TransactionType != WITHDRAWAL || t.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type cann't be nothing but withdrawal / deposit ")
	}
	if t.Amount < 5000 {
		return errs.NewValidationError("Transaction amount cann't be less than 5000")
	}
	return nil
}
func (t TransactionRequest) TransactionTypeIsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
