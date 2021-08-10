package account

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetAccountsResponse struct {
	Accounts []entries.Account `json:"accounts"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}
