package transfer

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) ListTransfers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	header := r.Header.Get("Authorization")

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "listTransfers"),
	)

	token, err := CheckAuthorizationHeaderType(header)
	e := errorStruct{l: l, token: token, w: w}
	if err != nil {
		l.With(zap.Any("type", http.StatusBadRequest)).Error(err.Error())
		e.errorCreate(err)
		return
	}

	Transfers, err := s.transfer.GetTransfers(token)
	if err != nil {
		e.errorList(err)
		return
	}

	response := GetTransfersResponse{Transfers: Transfers}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain.ErrInvalidToken):
		e.w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain.ErrInvalidID):
		e.w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain.ErrParse) || errors.Is(err, ErrInvalidCredential):
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to list transfers", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
