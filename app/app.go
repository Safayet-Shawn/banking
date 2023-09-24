package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Safayet-Shawn/banking/domain"
	"github.com/Safayet-Shawn/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASS") == "" ||
		os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" ||
		os.Getenv("DB_DRIVER") == "" {
		log.Fatal(" Environment veriable not defined !!!")
	}
}
func Start() {
	//define route
	// mux := http.NewServeMux()
	sanityCheck()
	router := mux.NewRouter()
	//wiring the application
	dbClient := getDbClient()
	NewCustomerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	NewAccountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandler{service.NewCustomerService(NewCustomerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(NewAccountRepositoryDb)}
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ah.NewAcount).Methods(http.MethodPost)
	// router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//starting server
	Host := os.Getenv("Server_Add")
	Port := os.Getenv("Server_Port")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", Host, Port), router))

}
func getDbClient() *sqlx.DB {
	Dbuser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbDriver := os.Getenv("DB_DRIVER")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Dbuser, DbPass, DbHost, DbPort, DbName)
	client, err := sqlx.Open(DbDriver, dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
