package account

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	accounts, err := s.account.GetAccounts()
	response := GetAccountsResponse{Accounts: accounts}
	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleAccounts",
	})
	e := errorStruct{l: l, w: w}
	if err != nil {
		e.errorList(err)
		return
	}
	l.WithFields(log.Fields{
		"type": http.StatusOK,
	}).Info("list the accounts successfully!")
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	if err.Error() == domain2.ErrInsert.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
		}).Error(err)
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	} else {
		e.w.WriteHeader(http.StatusBadRequest)
	}
}
