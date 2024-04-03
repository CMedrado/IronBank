package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	accounts, err := s.account.GetAccounts()
	response := GetAccountsResponse{Accounts: accounts}
	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "listAccounts"),
	)
	e := errorStruct{l: l, w: w}
	if err != nil {
		e.errorList(err)
		return
	}

	l.With(zap.Any("type", http.StatusOK)).Info("list the accounts successfully!")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain2.ErrInsert):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to list accounts", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
