package https

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (s *ServerAccount) processAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processAccount",
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idAccount, err := s.account.CreateAccount(requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given the balance amount is invalid":
			l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		case "given cpf is invalid":
			l.WithFields(log.Fields{
				"type": http.StatusNotAcceptable,
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := CreateResponse{ID: idAccount}

	l.WithFields(log.Fields{
		"type":       http.StatusCreated,
		"request_id": response,
	}).Info("account created successfully!")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) handleAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := GetAccountsResponse{Accounts: s.account.GetAccounts()}
	s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleAccounts",
		"type":   http.StatusOK,
	}).Info("accounts handled successfully!")
	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) handleBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	IntID, _ := strconv.Atoi(id)
	balance, err := s.account.GetBalance(IntID)
	w.Header().Set("content-type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleBalance",
	})

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given id is invalid":
			l.WithFields(log.Fields{
				"type":       http.StatusNotAcceptable,
				"request_id": id,
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	l.WithFields(log.Fields{
		"type":       http.StatusOK,
		"request_id": id,
	}).Info("balance handled successfully!")
	response := BalanceResponse{Balance: balance}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
