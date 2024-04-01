package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

type TransferResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetTransfersResponse struct {
	Transfers []entities.Transfer `json:"transfers"`
}

type GetAccountsResponse struct {
	Accounts []entities.Account `json:"accounts"`
}

type GetStatisticTransfersResponse struct {
	Transfers int64 `json:"transfers_count"`
}

type GetRankTransfersResponse struct {
	Transfers []string `json:"transfers_rank"`
}
