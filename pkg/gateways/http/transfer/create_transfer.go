package transfer

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	"github.com/CMedrado/DesafioStone/pkg/domain/transfer"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
)

func (s *Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var requestBody TransfersRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	header := r.Header.Get("Authorization")

	token, err := CheckAuthorizationHeaderType(header)

	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "createTransfer"),
	)
	e := errorStruct{l: l, token: token, w: w}

	if err != nil {
		e.errorCreate(err)
		l.Error("error check autorization header type", zap.Error(err))
		return
	}

	accountOriginID, tokenOriginID, err := authentication.DecoderToken(token)
	if err != nil {
		e.errorCreate(err)
		l.Error("error decoder token", zap.Error(err))
		return
	}

	accountOrigin, err := s.account.SearchAccount(accountOriginID)
	if err != nil {
		e.errorCreate(err)
		l.Error("error search account, account id", zap.Error(err))
		return
	}

	accountToken, err := s.login.GetTokenID(tokenOriginID)
	if err != nil {
		e.errorCreate(err)
		l.Error("error get token id", zap.Error(err))
		return
	}

	accountDestinationIdUUID, err := uuid.Parse(requestBody.AccountDestinationID)
	if err != nil {
		e.errorCreate(err)
		l.Error("error parse", zap.Error(err))
		return
	}

	accountDestination, err := s.account.SearchAccount(accountDestinationIdUUID)
	if err != nil {
		e.errorCreate(err)
		l.Error("error search account, account destination id", zap.Error(err))
		return
	}

	err, id, accountOrigin, accountDestination := s.transfer.CreateTransfers(r.Context(), accountOriginID, accountToken, token, accountOrigin, accountDestination, requestBody.Amount, accountDestinationIdUUID)
	if err != nil {
		e.errorCreate(err)
		return
	}

	err = s.account.UpdateBalance(accountOrigin, accountDestination)
	if err != nil {
		e.errorCreate(err)
		l.Error("error update balance", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	l.With(zap.Any("type", http.StatusOK)).Info("create transfer successfully!")

	response := TransferResponse{ID: id}
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(response)
}

type errorStruct struct {
	l     *zap.Logger
	token string
	w     http.ResponseWriter
}

func (e errorStruct) errorCreate(err error) {
	ErrJson := http2.ErrorsResponse{Errors: err.Error()}
	switch {
	case errors.Is(err, transfer.ErrWithoutBalance) ||
		errors.Is(err, transfer.ErrInvalidAmount) ||
		errors.Is(err, transfer.ErrSameAccount) ||
		errors.Is(err, domain2.ErrParse) ||
		errors.Is(err, ErrInvalidCredential):
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain2.ErrInvalidToken):
		e.w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain2.ErrInvalidID) ||
		errors.Is(err, transfer.ErrInvalidDestinationID):
		e.w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	case errors.Is(err, domain2.ErrInsert) ||
		errors.Is(err, domain2.ErrSelect):
		e.w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	default:
		e.l.Error("failed to create transfer", zap.Error(err))
		e.w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(e.w).Encode(ErrJson)
	}
}
