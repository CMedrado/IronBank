package https

import (
	"encoding/json"
	"net/http"
)

func (s *ServerAccount) handleTransfers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization:")

	Transfers, err := accountUseCase.GetTransfers(token)
	if err != nil {
		switch err.Error() {
		case "given id is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
		case "given token is invalid":
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Transfers)
}

func (s *ServerAccount) processTransfer(w http.ResponseWriter, r *http.Request) {
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	token := r.Header.Get("Authorization:")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, id := accountUseCase.CreateTransfers(token, requestBody.AccountDestinationID, requestBody.Amount)

	if err != nil {
		switch err.Error() {
		case "given id is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
		case "account without balance":
			w.WriteHeader(http.StatusPaymentRequired)
		case "given token is invalid":
			w.WriteHeader(http.StatusUnauthorized)
		case "given amount is invalid":
			w.WriteHeader(http.StatusPaymentRequired)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := TransferResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}
