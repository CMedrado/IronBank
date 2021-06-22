package transfer

import (
	"github.com/CMedrado/DesafioStone/storage/file/transfer"
)

type Repository interface {
	ReturnTransfers() []transfer.Transfer
	SaveTransfers(transfer transfer.Transfer)
}
