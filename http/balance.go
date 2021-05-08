package http

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/gorilla/mux.v1.8.0"
	"net/http"
	"strconv"
)

func (s *ServerAccount) GetBalance(w http.ResponseWriter, r *http.Request) {
	aux := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(aux)
	balance := domain.GetBalance(id)
	json.NewEncoder(w).Encode(balance)
}
