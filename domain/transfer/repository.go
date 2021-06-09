package transfer

import "github.com/CMedrado/DesafioStone/store"

type TransferRepository interface {
	GetTransfers(string) ([]store.Transfer, error)
	CreateTransfers(string, int, int) (error, int)
}
