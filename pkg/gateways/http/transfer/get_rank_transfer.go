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

func (s *Handler) GetRankTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "getRankTransfer"),
	)
	e := errorStruct{l: l, w: w}

	transfers, err := s.transfer.GetRankTransfer(r.Context())
	if err != nil {
		e.errorGetRankTransfer(err)
		return
	}

	response := GetRankTransfersResponse{Transfers: transfers}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorGetRankTransfer(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain.ErrGetRedis):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to get rank transfer", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
