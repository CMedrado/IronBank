package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetAccountsResponse struct {
	Accounts []domain.Account `json:"accounts"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}
