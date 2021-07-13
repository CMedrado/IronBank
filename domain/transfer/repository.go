package transfer

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/transfer"
)

type Repository interface {
	ReturnTransfer() ([]transfer.Transfer, error)
	SaveTransfer(transfer transfer.Transfer) error
}
