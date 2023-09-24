package app

import (
	"encoding/json"
	"net/http"

	"github.com/Safayet-Shawn/banking/dto"
	"github.com/Safayet-Shawn/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) NewAcount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		account, Apperror := ah.service.NewAccount(request)
		if Apperror != nil {
			WriteResponse(w, Apperror.Code, Apperror.Message)
		} else {
			WriteResponse(w, http.StatusCreated, account)
		}
	}
}
