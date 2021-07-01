package transfer

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) ListTransfers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	Transfers, err := s.transfer.GetTransfers(token)
	w.Header().Set("content-type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleTransfers",
	})

	if err != nil {
		ErrJson := https.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain.ErrInvalidToken.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusUnauthorized,
				"time": domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrParse.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
				"time": domain.CreatedAt(),
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
		"time": domain.CreatedAt(),
	}).Info("transfers handled successfully!")
	response := GetTransfersResponse{Transfers: Transfers}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
