package domain

import "github.com/Safayet-Shawn/banking/dto"

const WITHDRAWAL = "withdrawl"

type Transaction struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
	TransactionType string  `json:"transaction_type"`
}

func (t Transaction) IsWithwrals() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	} else {
		return false
	}
}
func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}

}
