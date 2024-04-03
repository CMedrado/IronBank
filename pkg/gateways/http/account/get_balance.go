package account

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	"github.com/gorilla/mux"
)

func (s *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	balance, err := s.account.GetBalance(id)
	w.Header().Set("content-type", "application/json")
	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "getBalance"),
	)
	e := errorStruct{l: l, w: w, id: id}
	if err != nil {
		e.errorBalance(err)
		return
	}
	l.With(zap.Any("type", http.StatusOK), zap.String("request_id", id)).Info("get balance successfully!")
	response := BalanceResponse{Balance: balance}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorBalance(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case err.Error() == domain2.ErrInvalidID.Error():
		e.w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case err.Error() == domain2.ErrSelect.Error():
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case err.Error() == domain2.ErrParse.Error():
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to get balance", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
