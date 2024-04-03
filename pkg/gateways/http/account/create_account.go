package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/account"
	http_server "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var requestBody CreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "createAccount"),
	)
	idAccount, err := s.account.CreateAccount(r.Context(), requestBody.Name, requestBody.CPF, requestBody.Secret, requestBody.Balance)
	w.Header().Set("Content-Type", "application/json")
	e := errorStruct{l: l, w: w}
	if err != nil {
		e.errorCreate(err)
		return
	}

	response := CreateResponse{ID: idAccount}

	l.With(zap.Any("request_id", response)).Info("account created successfully!")

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(response)
}

type errorStruct struct {
	l  *zap.Logger
	w  http.ResponseWriter
	id string
}

func (e errorStruct) errorCreate(err error) {
	ErrJson := http_server.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, account.ErrAccountExists) ||
		errors.Is(err, account.ErrBalanceAbsent) ||
		errors.Is(err, domain.ErrInvalidCPF):
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain.ErrInsert) ||
		errors.Is(err, domain.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to create account", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
