package account

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/account"
	http_server "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processAccount",
	})
	idAccount, err := s.account.CreateAccount(r.Context(), requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)
	w.Header().Set("Content-Type", "application/json")
	e := errorStruct{l: l, w: w}
	if err != nil {
		e.errorCreate(err)
		return
	}

	response := CreateResponse{ID: idAccount}

	l.WithFields(log.Fields{
		"type":       http.StatusCreated,
		"request_id": response,
	}).Info("account created successfully!")

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(response)
}

type errorStruct struct {
	l  *log.Entry
	w  http.ResponseWriter
	id string
}

func (e errorStruct) errorCreate(err error) {
	ErrJson := http_server.ErrorsResponse{Errors: err.Error()}
	if err.Error() == account.ErrAccountExists.Error() ||
		err.Error() == account.ErrBalanceAbsent.Error() ||
		err.Error() == domain.ErrInvalidCPF.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusBadRequest,
		}).Error(err)
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	} else if err.Error() == domain.ErrInsert.Error() ||
		err.Error() == domain.ErrSelect.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
		}).Error(err)
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	} else {
		e.w.WriteHeader(http.StatusBadRequest)
	}
}
