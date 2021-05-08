package http

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/gorilla/mux.v1.8.0"
	"net/http"
	"strconv"
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

func (s *ServerAccount) GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(domain.GetAccounts())
}

func (s *ServerAccount) GetBalance(w http.ResponseWriter, r *http.Request) {
	aux := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(aux)
	balance := domain.GetBalance(id)
	json.NewEncoder(w).Encode(balance)
}
