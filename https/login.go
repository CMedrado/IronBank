package https

import (
	"encoding/json"
	"github.com/CMedrado/DesafioStone/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *ServerAccount) processLogin(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, token := s.login.AuthenticatedLogin(requestBody.CPF, requestBody.Secret)
	w.Header().Set("Content-Type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processLogin",
	})

	if err != nil {
		ErrJson := ErrorsResponse{Errors: err.Error()}
		switch err.Error() {
		case "given secret or CPF are incorrect":
			l.WithFields(log.Fields{
				"type": http.StatusUnauthorized,
				"time": domain.CreatedAt(),
			}).Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrJson)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	l.WithFields(log.Fields{
		"type":           http.StatusOK,
		"time":           domain.CreatedAt(),
		"response_token": token,
	}).Info("balance handled sucessfully!")

	response := TokenResponse{Token: token}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
