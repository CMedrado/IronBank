package account

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetAccountsResponse struct {
	Accounts []entities.Account `json:"accounts"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}
