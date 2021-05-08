package http

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	"net/http"
)

func (s *ServerAccount) GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(domain.GetAccounts())
}
