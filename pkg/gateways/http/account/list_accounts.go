package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "listAccounts"),
	)
	e := errorStruct{l: l, w: w}

	accounts, err := s.account.GetAccounts()
	if err != nil {
		e.errorList(err)
		return
	}

	response := GetAccountsResponse{Accounts: accounts}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to list accounts", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
