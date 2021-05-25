package https

import (
	"encoding/json"
	"net/http"
)

func (s *ServerAccount) processLogin(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, token := accountUseCase.AuthenticatedLogin(requestBody.CPF, requestBody.Secret)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given secret or CPF are incorrect":
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response := TokenResponse{Token: token}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
