package account

import (
	"encoding/json"
	"errors"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
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
		"time":       domain2.CreatedAt(),
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
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	if errors.Is(err, domain2.ErrAccountExists) || errors.Is(err, domain2.ErrInsert) || errors.Is(err, domain2.ErrSelect) || errors.Is(err, domain2.ErrBalanceAbsent) {
		e.l.WithFields(log.Fields{
			"type": http.StatusBadRequest,
			"time": domain2.CreatedAt(),
		}).Error(err)
		e.w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(e.w).Encode(ErrJson)
	} else if err.Error() == domain2.ErrInvalidCPF.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusNotAcceptable,
			"time": domain2.CreatedAt(),
		}).Error(err)
		e.w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(e.w).Encode(ErrJson)
	} else {
		e.w.WriteHeader(http.StatusBadRequest)
	}
	return
}
