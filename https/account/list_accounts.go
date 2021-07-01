package account

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
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
