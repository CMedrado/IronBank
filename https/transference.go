package https

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *ServerAccount) handleTransfers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	Transfers, err := s.transfer.GetTransfers(token)
	w.Header().Set("content-type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleTransfers",
	})

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
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

func (s *ServerAccount) processTransfer(w http.ResponseWriter, r *http.Request) {
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	token := r.Header.Get("Authorization")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processTransfer",
	})

	err, id := s.transfer.CreateTransfers(token, requestBody.AccountDestinationID, requestBody.Amount)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain.ErrInvalidDestinationID.Error() {
			l.WithFields(log.Fields{
				"type":          http.StatusNotAcceptable,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrWithoutBalance.Error() {
			l.WithFields(log.Fields{
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrInvalidToken.Error() {
			l.WithFields(log.Fields{
				"type": http.StatusUnauthorized,
				"time": domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrInvalidAmount.Error() {
			l.WithFields(log.Fields{
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrSameAccount.Error() {
			l.WithFields(log.Fields{
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else if err.Error() == domain.ErrParse.Error() {
			l.WithFields(log.Fields{
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	l.WithFields(log.Fields{
		"type":       http.StatusCreated,
		"time":       domain.CreatedAt(),
		"request_id": id,
	}).Info("create transfer successfully!")

	response := TransferResponse{ID: id}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
