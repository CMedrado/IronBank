package transfer

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) GetStatisticTransfers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	transfers, err := s.transfer.GetStatisticTransfer(r.Context())
	response := GetStatisticTransfersResponse{Transfers: transfers}
	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "getStatisticTransfer",
	})
	e := errorStruct{l: l, w: w}
	if err != nil {
		e.errorGetStatisticTranfer(err)
		return
	}
	l.WithFields(log.Fields{
		"type": http.StatusOK,
	}).Info("get the statistic transfer successfully!")
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorGetStatisticTranfer(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	if err.Error() == domain2.ErrGetRedis.Error() {
		e.l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
		}).Error(err)
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	} else {
		e.w.WriteHeader(http.StatusBadRequest)
	}
}
