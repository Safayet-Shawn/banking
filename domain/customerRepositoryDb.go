package domain

import (
	"log"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

// receiver func
func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllsql := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers"
	rows, err := d.client.Query(findAllsql)
	if err != nil {
		log.Println("Error while querying customer table where err is ", err.Error)
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scaning customer table where err is ", err.Error)
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, error) {
	FindCustomer := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers WHERE customer_id =?"
	row := d.client.QueryRow(FindCustomer, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		log.Println("Error while quering customer table where id and err respectively", id, err)
		return nil, err
	}
	return &c, nil
}

// all new is helper function
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:itsshawn@007@@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
