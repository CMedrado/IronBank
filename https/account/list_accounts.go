package account

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	if err != nil {
		ErrJson := https.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain.ErrUpdate.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	l.WithFields(log.Fields{
		"type": http.StatusOK,
	}).Info("account created successfully!")
	json.NewEncoder(w).Encode(response)
}
