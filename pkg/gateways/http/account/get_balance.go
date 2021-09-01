package account

import (
	"encoding/json"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	balance, err := s.account.GetBalance(id)
	w.Header().Set("content-type", "application/json")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "handleBalance",
	})
	e := errorStruct{l: l, w: w, id: id}
	if err != nil {
		e.errorBalance(err)
		return
	}
	l.WithFields(log.Fields{
		"type":       http.StatusOK,
		"request_id": id,
	}).Info("balance handled successfully!")
	response := BalanceResponse{Balance: balance}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorBalance(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	if err.Error() == domain2.ErrInvalidID.Error() {
		e.l.WithFields(log.Fields{
			"type":       http.StatusNotFound,
			"request_id": e.id,
			"time":       domain2.CreatedAt(),
		}).Error(err)
		e.w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(e.w).Encode(ErrJson)
	} else if err.Error() == domain2.ErrSelect.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
			"time": domain2.CreatedAt(),
		}).Error(err)
		e.w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(e.w).Encode(ErrJson)
	} else if err.Error() == domain2.ErrParse.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusBadRequest,
			"time": domain2.CreatedAt(),
		}).Error(err)
		e.w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(e.w).Encode(ErrJson)
	} else {
		e.w.WriteHeader(http.StatusBadRequest)
	}
}
