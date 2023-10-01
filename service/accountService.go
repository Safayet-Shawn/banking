package service

import (
	"time"

	"github.com/Safayet-Shawn/banking/domain"
	"github.com/Safayet-Shawn/banking/dto"
	"github.com/Safayet-Shawn/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.Apperror)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.Apperror)
}
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.Apperror) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := d.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResonseDto()
	return &response, nil
}
func (d DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.Apperror) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.TransactionTypeIsWithdrawal() {
		account, err := d.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient amount found in the account")
		}
	}
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02"),
	}
	transaction, appErr := d.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, errs.NewUnexpectedServerError("Failed to get account info from seever")
	}
	response := transaction.ToDto()
	return &response, nil
}
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
