package domain

import "github.com/CMedrado/DesafioStone/store"

type LoginRepository interface {
	AuthenticatedLogin(string, string) (error, string)
}

type TransferRepository interface {
	GetTransfers(string) ([]store.Transfer, error)
	CreateTransfers(string, int, int) (error, int)
}
