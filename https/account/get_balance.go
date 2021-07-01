package account

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	balance, err := s.account.GetBalance(id)
	w.Header().Set("content-type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleBalance",
	})

	if err != nil {
		ErrJson := https.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain.ErrInvalidID.Error() {
			l.WithFields(log.Fields{
				"type":       http.StatusNotAcceptable,
				"request_id": id,
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrParse.Error() {
			l.WithFields(log.Fields{
				"type":       http.StatusBadRequest,
				"request_id": id,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else {
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
