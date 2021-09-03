package transfer

import (
	"encoding/json"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	"github.com/CMedrado/DesafioStone/pkg/domain/transfer"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	header := r.Header.Get("Authorization")

	token, err := CheckCredential(header)

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processTransfer",
		"token":  token,
	})
	e := errorStruct{l: l, token: token, w: w}

	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "checkcredential",
		}).Error(err)
		e.errorCreate(err)
		return
	}

	accountOriginID, tokenOriginID, err := authentication.DecoderToken(token)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusUnauthorized,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "decoderToken",
		}).Error(err)
		e.errorCreate(err)
		return
	}
	accountOrigin, err := s.account.SearchAccount(accountOriginID)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusInternalServerError,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "searchAccount",
		}).Error(err)
		e.errorCreate(err)
		return
	}
	accountToken, err := s.login.GetTokenID(tokenOriginID)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusInternalServerError,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "getTokenID",
		}).Error(err)
		e.errorCreate(err)
		return
	}
	accountDestinationIdUUID, err := uuid.Parse(requestBody.AccountDestinationID)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "Parse",
		}).Error(err)
		e.errorCreate(err)
		return
	}
	accountDestination, err := s.account.SearchAccount(accountDestinationIdUUID)
	if err != nil {
		l.WithFields(log.Fields{
			"type":  http.StatusInternalServerError,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "searchAccount",
		}).Error(err)
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
		l.WithFields(log.Fields{
			"type":  http.StatusInternalServerError,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "searchAccount",
		}).Error(err)
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
		if err.Error() == transfer.ErrWithoutBalance.Error() ||
			err.Error() == transfer.ErrInvalidAmount.Error() ||
			err.Error() == transfer.ErrSameAccount.Error() ||
			err.Error() == domain2.ErrParse.Error() ||
			err.Error() == ErrInvalidCredential.Error() {
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
		} else if err.Error() == domain2.ErrInvalidID.Error() ||
			err.Error() == transfer.ErrInvalidDestinationID.Error() {
			e.l.WithFields(log.Fields{
				"type":          http.StatusNotFound,
				"time":          domain2.CreatedAt(),
				"request_token": e.token,
			}).Error(err)
			e.w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrInsert.Error() ||
			err.Error() == domain2.ErrSelect.Error() {
			e.l.WithFields(log.Fields{
				"type":          http.StatusInternalServerError,
				"time":          domain2.CreatedAt(),
				"request_token": e.token,
			}).Error(err)
			e.w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else {
			e.w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
}
