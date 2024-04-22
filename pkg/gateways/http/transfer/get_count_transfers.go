package transfer

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) GetCountTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "getCountTransfer"),
	)

	transfers, err := s.transfer.GetCountTransfer(r.Context())
	if err != nil {
		ErrJson := http2.ErrorsResponse{Errors: err.Error()}
		l.Error("failed to get rank transfer", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson)
		return
	}

	response := GetCountTransferResponse{Transfers: transfers}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
