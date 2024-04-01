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
	SearchAccount(id uuid.UUID) (entities.Account, error)
	UpdateBalance(accountOrigin entities.Account, accountDestination entities.Account) error
	GetAccountCPF(ctx context.Context, cpf string) (entities.Account, error)
}

type LoginUseCase interface {
	AuthenticatedLogin(secret string, id entities.Account) (error, string)
	GetTokenID(id uuid.UUID) (entities.Token, error)
}

type TransferUseCase interface {
	GetTransfers(accountOrigin entities.Account, accountToken entities.Token, token string) ([]entities.Transfer, error)
	CreateTransfers(ctx context.Context, accountOriginID uuid.UUID, accountToken entities.Token, token string, accountOrigin entities.Account, accountDestination entities.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entities.Account, entities.Account)
	GetStatisticTransfer(ctx context.Context) (int64, error)
	GetRankTransfer(ctx context.Context) ([]string, error)
}
