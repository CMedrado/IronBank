package domain

import "github.com/google/uuid"

type AccountUseCase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (uuid.UUID, error)
	GetBalance(id string) (int, error)
	GetAccounts() []Account
	SearchAccount(id uuid.UUID) Account
	UpdateBalance(accountOrigin Account, accountDestination Account)
	GetAccountCPF(cpf string) Account
}

type LoginUseCase interface {
	AuthenticatedLogin(cpf, secret string) (error, string)
	GetTokenID(id uuid.UUID) Token
}

type TransferUseCase interface {
	GetTransfers(token string) ([]Transfer, error)
	CreateTransfers(token string, accountDestinationID string, amount int) (error, uuid.UUID)
}
