package domain

import (
	"fmt"

	"time"

	"github.com/Safayet-Shawn/banking/errs"
	"github.com/Safayet-Shawn/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// receiver func
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.Apperror) {
	customers := make([]Customer, 0)
	var err error
	if status == "" {
		findAllsql := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers"
		err = d.client.Select(&customers, findAllsql)
	} else {
		findAllsql := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers WHERE status= ?"
		err = d.client.Select(&customers, findAllsql, status)
	}
	if err != nil {
		msg := fmt.Sprintf("Error while querying customer table where err is %v", err)
		logger.Error(msg)
		return nil, errs.NewUnexpectedServerError("Error from database")
	}
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.Apperror) {
	var c Customer
	FindCustomer := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers WHERE customer_id =?"
	err := d.client.Get(&c, FindCustomer, id)
	if err != nil {
		meg := fmt.Sprintf("Error while quering customer table where id = %v and err = %v respectively", id, err)
		logger.Error(meg)
		return nil, errs.NewUnexpectedServerError(meg)
	}
	return &c, nil
}

// all new is helper function
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:itsshawn@007@@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
