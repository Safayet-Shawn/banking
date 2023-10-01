package app

import (
	"encoding/json"
	"net/http"

	"github.com/Safayet-Shawn/banking/dto"
	"github.com/Safayet-Shawn/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) NewAcount(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	customerId := variable["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, Apperror := ah.service.NewAccount(request)
		if Apperror != nil {
			WriteResponse(w, Apperror.Code, Apperror.Message)
		} else {
			WriteResponse(w, http.StatusCreated, account)
		}
	}
}
func (ah *AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	//decoding incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
		account, appErr := ah.service.MakeTransaction(request)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, account)
		}
	}
}
