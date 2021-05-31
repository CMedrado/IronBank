package https

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
			log.WithFields(log.Fields{
				"module": "https",
				"method": "processAccount",
				"type":   http.StatusBadRequest,
				"time":   domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		case "given cpf is invalid":
			log.WithFields(log.Fields{
				"module": "https",
				"method": "processAccount",
				"type":   http.StatusNotAcceptable,
				"time":   domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := CreateResponse{ID: idAccount}

	log.WithFields(log.Fields{
		"module": "https",
		"method": "processAccount",
		"type":   http.StatusCreated,
		"id":     response,
		"time":   domain.CreatedAt(),
	}).Info("account created sucessfully!")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) handleAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := GetAccountsResponse{Accounts: accountUseCase.GetAccounts()}
	log.WithFields(log.Fields{
		"module": "https",
		"method": "handleAccounts",
		"type":   http.StatusOK,
		"time":   domain.CreatedAt(),
	}).Info("accounts handled sucessfully!")
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
			log.WithFields(log.Fields{
				"module":     "https",
				"method":     "handleBalance",
				"type":       http.StatusNotAcceptable,
				"request_id": id,
				"time":       domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	log.WithFields(log.Fields{
		"module":     "https",
		"method":     "handleBalance",
		"type":       http.StatusOK,
		"request_id": id,
		"time":       domain.CreatedAt(),
	}).Info("balance handled sucessfully!")
	response := BalanceResponse{Balance: balance}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
