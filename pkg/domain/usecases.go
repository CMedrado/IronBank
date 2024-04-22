package domain

import (
	"context"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

type AccountUseCase interface {
	CreateAccount(ctx context.Context, name string, cpf string, secret string, balance int) (uuid.UUID, error)
	GetBalance(id string) (int, error)
	GetAccounts() ([]entities.Account, error)
	UpdateBalance(accountOrigin entities.Account, accountDestination entities.Account) error
	GetAccountCPF(ctx context.Context, cpf string) (entities.Account, error)
	GetAccountID(id uuid.UUID) (entities.Account, error)
}

type LoginUseCase interface {
	AuthenticatedLogin(secret string, cpf string) (error, string)
	GetTokenID(id uuid.UUID) (entities.Token, error)
}

type TransferUseCase interface {
	GetTransfers(token string) ([]entities.Transfer, error)
	CreateTransfers(ctx context.Context, token string, amount int, accountDestinationIdUUID string) (error, uuid.UUID)
	GetCountTransfer(ctx context.Context) (int64, error)
	GetRankTransfer(ctx context.Context) ([]string, error)
}
