package authentication

import (
	"encoding/json"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processLogin",
	})
	e := errorStruct{l: l, w: w}
	err, cpf := domain2.CheckCPF(requestBody.CPF)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"where": "checkCPF",
		}).Error(err)
		e.errorLogin(err)
		return
	}
	account, err := s.account.GetAccountCPF(cpf)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"where": "getAccountCPF",
		}).Error(err)
		e.errorLogin(err)
		return
	}
	err, token := s.login.AuthenticatedLogin(requestBody.Secret, account)
	if err != nil {
		e.errorLogin(err)
		return
	}

	l.WithFields(log.Fields{
		"type":           http.StatusOK,
		"time":           domain2.CreatedAt(),
		"response_token": token,
	}).Info("sucessfully authentificated!")

	response := TokenResponse{Token: token}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

type errorStruct struct {
	l *log.Entry
	w http.ResponseWriter
}

func (e errorStruct) errorLogin(err error) {
	if err != nil {
		ErrJson := http2.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain2.ErrLogin.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusUnauthorized,
				"time": domain2.CreatedAt(),
			}).Error(err)
			e.w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrInsert.Error() || err.Error() == domain2.ErrSelect.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
				"time": domain2.CreatedAt(),
			}).Error(err)
			e.w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else {
			e.w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
}
