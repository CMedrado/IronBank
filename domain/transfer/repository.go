package transfer

import (
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
)

type Repository interface {
	GetTransfers(accountOriginID int) map[int]store_transfer.Transfer
	PostTransferID(transfer store_transfer.Transfer, id int)
}
