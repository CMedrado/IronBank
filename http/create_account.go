package http

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"net/http"
)

func (s *ServerAccount) CreatedAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody CreatedRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	idAccount, err := domain.CreatedAccount(requestBody.Name, requestBody.CPF, requestBody.Secret)

	json.NewEncoder(w).Encode(idAccount)
}
