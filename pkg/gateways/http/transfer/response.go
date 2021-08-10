package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type TransferResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetTransfersResponse struct {
	Transfers []entries.Transfer `json:"transfers"`
}

type GetAccountsResponse struct {
	Accounts []entries.Account `json:"accounts"`
}
