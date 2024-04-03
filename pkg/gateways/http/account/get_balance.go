package account

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	"github.com/gorilla/mux"
)

func (s *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "getBalance"),
	)
	e := errorStruct{l: l, w: w, id: id}

	balance, err := s.account.GetBalance(id)
	if err != nil {
		e.errorBalance(err)
		return
	}

	response := BalanceResponse{Balance: balance}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorBalance(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case err.Error() == domain.ErrInvalidID.Error():
		e.w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case err.Error() == domain.ErrSelect.Error():
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case err.Error() == domain.ErrParse.Error():
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to get balance", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
