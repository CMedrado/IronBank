package account

import (
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody GetRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "getAccount",
	})
	e := errorStruct{l: l, w: w, id: requestBody.CPF}
	if err != nil {
		e.errorBalance(err)
		return
	}

	err, cpf := domain.CheckCPF(requestBody.CPF)
	if err != nil {
		e.errorGet(domain.ErrInvalidCPF)
		return
	}
	account, err := s.account.GetAccountCPF(cpf)
	if err != nil {
		e.errorGet(err)
		return
	}

	l.WithFields(log.Fields{
		"type":       http.StatusOK,
		"request_id": requestBody.CPF,
	}).Info("get account successfully!")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(account)
}

func (e errorStruct) errorGet(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain.ErrInsert) || errors.Is(err, domain.ErrSelect):
		e.l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
		}).Error(err)
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
