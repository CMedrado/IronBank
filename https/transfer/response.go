package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type TransferResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetTransfersResponse struct {
	Transfers []domain.Transfer `json:"transfers"`
}

type GetAccountsResponse struct {
	Accounts []domain.Account `json:"accounts"`
}
