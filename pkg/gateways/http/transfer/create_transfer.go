package transfer

import (
	"encoding/json"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	token := r.Header.Get("Authorization")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processTransfer",
	})
	e := errorStruct{l: l, token: token, w: w}
	accountOriginID, tokenOriginID, err := DecoderToken(token)
	if err != nil {
		e.errorCreate(err)
		return
	}
	accountOrigin, err := s.account.SearchAccount(accountOriginID)
	if err != nil {
		e.errorCreate(err)
		return
	}
	accountToken, err := s.login.GetTokenID(tokenOriginID)
	if err != nil {
		e.errorCreate(err)
		return
	}
	accountDestinationIdUUID, err := uuid.Parse(requestBody.AccountDestinationID)
	if err != nil {
		e.errorCreate(err)
		return
	}
	accountDestination, err := s.account.SearchAccount(accountDestinationIdUUID)
	if err != nil {
		e.errorCreate(err)
		return
	}
	err, id, accountOrigin, accountDestination := s.transfer.CreateTransfers(accountOriginID, accountToken, token, accountOrigin, accountDestination, requestBody.Amount, accountDestinationIdUUID)
	if err != nil {
		e.errorCreate(err)
		return
	}
	err = s.account.UpdateBalance(accountOrigin, accountDestination)
	if err != nil {
		e.errorCreate(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	l.WithFields(log.Fields{
		"type":       http.StatusCreated,
		"time":       domain2.CreatedAt(),
		"request_id": id,
	}).Info("create transfer successfully!")

	response := TransferResponse{ID: id}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

type errorStruct struct {
	l     *log.Entry
	token string
	w     http.ResponseWriter
}

func (e errorStruct) errorCreate(err error) {
	if err != nil {
		ErrJson := http2.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain2.ErrWithoutBalance.Error() ||
			err.Error() == domain2.ErrInvalidAmount.Error() ||
			err.Error() == domain2.ErrSameAccount.Error() ||
			err.Error() == domain2.ErrInsert.Error() ||
			err.Error() == domain2.ErrParse.Error() ||
			err.Error() == domain2.ErrSelect.Error() ||
			err.Error() == domain2.ErrInvalidDestinationID.Error() {
			e.l.WithFields(log.Fields{
				"type":          http.StatusBadRequest,
				"time":          domain2.CreatedAt(),
				"request_token": e.token,
			}).Error(err)
			e.w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrInvalidToken.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusUnauthorized,
				"time": domain2.CreatedAt(),
			}).Error(err)
			e.w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrInvalidID.Error() {
			e.l.WithFields(log.Fields{
				"type":          http.StatusNotAcceptable,
				"time":          domain2.CreatedAt(),
				"request_token": e.token,
			}).Error(err)
			e.w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else {
			e.w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
}
