package transfer

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) ListTransfers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	header := r.Header.Get("Authorization")

	token, err := CheckAuthorizationHeaderType(header)

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processTransfer",
	})
	e := errorStruct{l: l, token: token, w: w}

	if err != nil {
		l.WithFields(log.Fields{
			"type": http.StatusBadRequest,
		}).Error(err)
		e.errorCreate(err)
		return
	}

	accountOriginID, tokenID, err := authentication.DecoderToken(token)
	if err != nil {
		e.errorList(err)
		return
	}
	accountOrigin, err := s.account.SearchAccount(accountOriginID)
	if err != nil {
		e.errorList(err)
		return
	}
	accountToken, err := s.login.GetTokenID(tokenID)
	if err != nil {
		e.errorList(err)
		return
	}
	Transfers, err := s.transfer.GetTransfers(accountOrigin, accountToken, token)
	if err != nil {
		e.errorList(err)
		return
	}
	l.WithFields(log.Fields{
		"type": http.StatusOK,
	}).Info("transfers handled successfully!")
	response := GetTransfersResponse{Transfers: Transfers}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	if err != nil {
		ErrJson := http2.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain2.ErrInvalidToken.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusUnauthorized,
			}).Error(err)
			e.w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrSelect.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusInternalServerError,
			}).Error(err)
			e.w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrInvalidID.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusNotFound,
			}).Error(err)
			e.w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrParse.Error() || err.Error() == ErrInvalidCredential.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
			}).Error(err)
			e.w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(e.w).Encode(ErrJson)
		} else {
			e.w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
}
