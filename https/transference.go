package https

import (
	"encoding/json"
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
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := GetTransfersResponse{Transfers: Transfers}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func (s *ServerAccount) processTransfer(w http.ResponseWriter, r *http.Request) {
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	token := r.Header.Get("Authorization")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, id := accountUseCase.CreateTransfers(token, requestBody.AccountDestinationID, requestBody.Amount)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given account destination id is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(ErrJson)
		case "given account without balance":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		case "given token is invalid":
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		case "given amount is invalid":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		case "given account is the same as the account destination":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := TransferResponse{ID: id}
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}
