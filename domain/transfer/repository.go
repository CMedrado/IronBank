package transfer

import "github.com/CMedrado/DesafioStone/store"

type Repository interface {
	GetTransfers(accountOriginID int) map[int]store.Transfer
	PostTransferID(transfer store.Transfer, id int)
}
