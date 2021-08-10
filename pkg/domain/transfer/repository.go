package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type Repository interface {
	ReturnTransfer(id uuid.UUID) ([]entries.Transfer, error)
	SaveTransfer(transfer entries.Transfer) error
}
