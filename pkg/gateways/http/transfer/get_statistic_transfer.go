package transfer

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) GetStatisticTransfers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	transfers, err := s.transfer.GetStatisticTransfer(r.Context())
	response := GetStatisticTransfersResponse{Transfers: transfers}
	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "getStatisticTransfer"),
	)
	e := errorStruct{l: l, w: w}
	if err != nil {
		e.errorGetStatisticTranfer(err)
		return
	}

	l.With(zap.Any("type", http.StatusOK)).Info("get statistic transfer successfully!")

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorGetStatisticTranfer(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, domain2.ErrGetRedis):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to get rank transfer", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
