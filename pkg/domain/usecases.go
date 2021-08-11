package domain

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

type AccountUseCase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (uuid.UUID, error)
	GetBalance(id string) (int, error)
	GetAccounts() ([]entities.Account, error)
	SearchAccount(id uuid.UUID) (entities.Account, error)
	UpdateBalance(accountOrigin entities.Account, accountDestination entities.Account) error
	GetAccountCPF(cpf string) (entities.Account, error)
}

type LoginUseCase interface {
	AuthenticatedLogin(secret string, id entities.Account) (error, string)
	GetTokenID(id uuid.UUID) (entities.Token, error)
}

type TransferUseCase interface {
	GetTransfers(accountOriginID uuid.UUID, accountToken entities.Token, token string) ([]entities.Transfer, error)
	CreateTransfers(accountOriginID uuid.UUID, accountToken entities.Token, token string, accountOrigin entities.Account, accountDestination entities.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entities.Account, entities.Account)
}
