package transfer

import "github.com/CMedrado/DesafioStone/store"

type TransferRepository interface {
	GetTransfers(accountOriginID int) map[int]store.Transfer
	PostTransferAccountID(transfer store.Transfer) store.StoredTransferID
	PostTransferID(transfer store.Transfer, id int)
}
