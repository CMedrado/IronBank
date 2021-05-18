package https

import (
	"encoding/json"
	"net/http"
)

func (s *ServerAccount) AuthenticatedLogin(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, token := accountUseCase.AuthenticatedLogin(requestBody.CPF, requestBody.Secret)

	if err != nil {
		switch err.Error() {
		case "given cpf is invalid":
			w.WriteHeader(http.StatusNotAcceptable)
		case "given secret is invalid":
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := TokenRequest{Token: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}
