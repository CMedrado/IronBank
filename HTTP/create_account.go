package HTTP

import (
	"encoding/json"
	"net/http"
)

func (s *ServidorConta) AçãoCriarConta(w http.ResponseWriter, r *http.Request) {
	var conta CreatedRequest
	err := json.NewDecoder(r.Body).Decode(&conta)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}