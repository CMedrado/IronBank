package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestCPF := mux.Vars(r)["cpf"]

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "getAccount"),
	)
	e := errorStruct{l: l, w: w, id: requestCPF}

	account, err := s.account.GetAccountCPF(r.Context(), requestCPF)
	if err != nil {
		e.errorGet(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(account)
}

func (e errorStruct) errorGet(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain.ErrInvalidCPF):
		e.w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain.ErrInsert) || errors.Is(err, domain.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to get account", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
