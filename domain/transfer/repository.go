package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
)

type Repository interface {
	ReturnTransfer() ([]domain.Transfer, error)
	SaveTransfer(transfer domain.Transfer) error
}
