package http

import (
	"encoding/json"
	"errors"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
	"github.com/gorilla/mux.v1.8.0"
	"net/http"
)

var (
	accountStorage = store.NewStoredAccount()
	accountUseCase = domain.AccountUsecase{Store: accountStorage}
)

func (s *ServerAccount) CreatedAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody CreatedRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idAccount, err := accountUseCase.CreateAccount(requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)

	if err != nil {
		switch err {
		case errors.New("given cpf is invalid"):
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := CreatedRequest{ID: idAccount}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accountUseCase.GetAccounts())
}

func (s *ServerAccount) GetBalance(w http.ResponseWriter, r *http.Request) {
	cpf := mux.Vars(r)["cpf"]
	response, err := accountUseCase.GetBalance(cpf)

	if err != nil {
		switch err {
		case errors.New("given account is invalid"):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
