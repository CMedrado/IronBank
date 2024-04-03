package authentication

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestBody LoginRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "processLogin"),
	)
	e := errorStruct{l: l, w: w}

	account, err := s.account.GetAccountCPF(r.Context(), requestBody.CPF)
	if err != nil {
		e.errorLogin(err)
		return
	}

	err, token := s.login.AuthenticatedLogin(requestBody.Secret, account)
	if err != nil {
		e.errorLogin(err)
		return
	}

	response := TokenResponse{Token: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		e.errorLogin(err)
		return
	}
}

type errorStruct struct {
	l *zap.Logger
	w http.ResponseWriter
}

func (e errorStruct) errorLogin(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, authentication.ErrLogin):
		e.w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain.ErrInsert) || errors.Is(err, domain.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to login", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
