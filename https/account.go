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
	accountTransfer = store.NewStoredTransferID()
	accountToken    = store.NewStoredToked()
	accountLogin    = store.NewStoredLogin()
	accountUseCase  = domain.AccountUseCase{accountStorage, accountLogin, accountToken, accountTransfer}
)

func (s *ServerAccount) processAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody CreatedRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idAccount, err := accountUseCase.CreateAccount(requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)

	if err != nil {
		switch err.Error() {
		case "given cpf is invalid":
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

func (s *ServerAccount) handleAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(accountUseCase.GetAccounts())
}

func (s *ServerAccount) handleBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	IntID, _ := strconv.Atoi(id)
	response, err := accountUseCase.GetBalance(IntID)

	if err != nil {
		switch err.Error() {
		case "given account is invalid":
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}