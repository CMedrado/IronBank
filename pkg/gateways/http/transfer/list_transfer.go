package transfer

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) ListTransfers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("content-type", "application/json")
	header := r.Header.Get("Authorization")

	token, err := CheckAuthorizationHeaderType(header)

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "listTransfers"),
	)
	e := errorStruct{l: l, token: token, w: w}

	if err != nil {
		l.With(zap.Any("type", http.StatusBadRequest)).Error(err.Error())
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

	l.With(zap.Any("type", http.StatusOK)).Info("list transfers successfully!")
	response := GetTransfersResponse{Transfers: Transfers}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain2.ErrInvalidToken):
		e.w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain2.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain2.ErrInvalidID):
		e.w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain2.ErrParse) || errors.Is(err, ErrInvalidCredential):
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to list transfers", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
