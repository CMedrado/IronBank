package account

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/account"
	http_server "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	idAccount, err := s.account.CreateAccount(requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)
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
		"time":       domain.CreatedAt(),
	}).Info("account created successfully!")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

type errorStruct struct {
	l  *log.Entry
	w  http.ResponseWriter
	id string
}

func (e errorStruct) errorCreate(err error) {
	ErrJson := http_server.ErrorsResponse{Errors: err.Error()}
	if err.Error() == account.ErrAccountExists.Error() ||
		err.Error() == domain.ErrInsert.Error() ||
		err.Error() == domain.ErrSelect.Error() ||
		err.Error() == account.ErrBalanceAbsent.Error() ||
		err.Error() == domain.ErrInvalidCPF.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusBadRequest,
			"time": domain.CreatedAt(),
		}).Error(err)
		e.w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(e.w).Encode(ErrJson)
	} else {
		e.w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
