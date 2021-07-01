package account

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
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
		ErrJson := https.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain.ErrAccountExists.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrBalanceAbsent.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrInvalidCPF.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusNotAcceptable,
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		} else {
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
