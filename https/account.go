package https

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	accountStorage  = store.NewStoredAccount()
	accountTransfer = store.NewStoredTransferAccountID()
	accountToken    = store.NewStoredToked()
	accountLogin    = store.NewStoredLogin()
	accountUseCase  = domain.AccountUseCase{Store: accountStorage, Login: accountLogin, Token: accountToken, Transfer: accountTransfer}
)

func (s *ServerAccount) processAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idAccount, err := accountUseCase.CreateAccount(requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given the balance amount is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		case "given cpf is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := CreateResponse{ID: idAccount}

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) handleAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	response := GetAccountsResponse{Accounts: accountUseCase.GetAccounts()}
	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) handleBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	IntID, _ := strconv.Atoi(id)
	balance, err := accountUseCase.GetBalance(IntID)
	w.Header().Set("content-type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given id is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	response := BalanceResponse{Balance: balance}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}
