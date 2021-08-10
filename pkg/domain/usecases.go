package domain

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type AccountUseCase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (uuid.UUID, error)
	GetBalance(id string) (int, error)
	GetAccounts() ([]entries.Account, error)
	SearchAccount(id uuid.UUID) (entries.Account, error)
	UpdateBalance(accountOrigin entries.Account, accountDestination entries.Account) error
	GetAccountCPF(cpf string) (entries.Account, error)
}

type LoginUseCase interface {
	AuthenticatedLogin(cpf, secret string, id entries.Account) (error, string)
	GetTokenID(id uuid.UUID) (entries.Token, error)
}

type TransferUseCase interface {
	GetTransfers(accountOriginID uuid.UUID, accountToken entries.Token, token string) ([]entries.Transfer, error)
	CreateTransfers(accountOriginID uuid.UUID, accountToken entries.Token, token string, accountOrigin entries.Account, accountDestination entries.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entries.Account, entries.Account)
}
