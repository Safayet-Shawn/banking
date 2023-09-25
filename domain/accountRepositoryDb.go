package domain

import (
	"fmt"
	"strconv"

	"github.com/Safayet-Shawn/banking/errs"
	"github.com/Safayet-Shawn/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.Apperror) {
	sqlInsert := "INSERT INTO Accounts(customer_id,opening_date,account_type,amount,status) values(?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		msg := fmt.Sprintf("Error while creatting new account value reason: ", err.Error())
		logger.Error(msg)
		return nil, errs.NewUnexpectedServerError("Unexpect Error while inserting value from server")
	}
	id, err := result.LastInsertId()
	if err != nil {
		msg := fmt.Sprintf("Error while creatting new account value reason: ", err.Error())
		logger.Error(msg)
		return nil, errs.NewUnexpectedServerError("Unexpect Error while inserting value from server")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.Apperror) {
	var a Account
	FindAccount := "SELECT * from account WHERE account_id=?"
	err := d.client.Get(&a, FindAccount, accountId)
	if err != nil {
		msg := fmt.Sprintf("Failed to find account info where account id is %v", accountId)
		logger.Error(msg)
		return nil, errs.NewUnexpectedServerError(msg)
	}
	return &a, nil
}
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{dbClient}
}
