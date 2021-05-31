package https

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *ServerAccount) handleTransfers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	Transfers, err := accountUseCase.GetTransfers(token)
	w.Header().Set("content-type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given token is invalid":
			log.WithFields(log.Fields{
				"module": "https",
				"method": "handleTransfers",
				"type":   http.StatusUnauthorized,
				"time":   domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	log.WithFields(log.Fields{
		"module": "https",
		"method": "handleTransfers",
		"type":   http.StatusOK,
		"time":   domain.CreatedAt(),
	}).Info("balance handled sucessfully!")
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

	err, id := accountUseCase.CreateTransfers(token, requestBody.AccountDestinationID, requestBody.Amount)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given account destination id is invalid":
			log.WithFields(log.Fields{
				"module":        "https",
				"method":        "processTransfer",
				"type":          http.StatusNotAcceptable,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		case "given account without balance":
			log.WithFields(log.Fields{
				"module":        "https",
				"method":        "processTransfer",
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		case "given token is invalid":
			log.WithFields(log.Fields{
				"module": "https",
				"method": "processTransfer",
				"type":   http.StatusUnauthorized,
				"time":   domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		case "given amount is invalid":
			log.WithFields(log.Fields{
				"module":        "https",
				"method":        "processTransfer",
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		case "given account is the same as the account destination":
			log.WithFields(log.Fields{
				"module":        "https",
				"method":        "processTransfer",
				"type":          http.StatusBadRequest,
				"time":          domain.CreatedAt(),
				"request_token": token,
			}).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	log.WithFields(log.Fields{
		"module":     "https",
		"method":     "handleTransfers",
		"type":       http.StatusCreated,
		"time":       domain.CreatedAt(),
		"request_id": id,
	}).Info("create transfer sucessfully!")

	response := TransferResponse{ID: id}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
