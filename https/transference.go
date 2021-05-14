package https

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (s *ServerAccount) GetTransfers(w http.ResponseWriter, r *http.Request) {
	var requestBody TokenRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Transfers, err := accountUseCase.GetTransfers(requestBody.AccountOriginID, requestBody.Token)

	if err != nil {
		switch err {
		case errors.New("given id is invalid"):
			w.WriteHeader(http.StatusNotAcceptable)
		case errors.New("given token is invalid"):
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Transfers)
}

func (s *ServerAccount) MakeTransfers(w http.ResponseWriter, r *http.Request) {
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, id := accountUseCase.MakeTransfers(requestBody.AccountOriginID, requestBody.Token, requestBody.AccountDestinationID, requestBody.Amount)

	if err != nil {
		switch err {
		case errors.New("given id is invalid"):
			w.WriteHeader(http.StatusNotAcceptable)
		case errors.New("account without balance"):
			w.WriteHeader(http.StatusPaymentRequired)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := TransfersRequest{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}
