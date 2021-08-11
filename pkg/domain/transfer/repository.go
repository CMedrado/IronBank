package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

type Repository interface {
	ReturnTransfer(id uuid.UUID) ([]entities.Transfer, error)
	SaveTransfer(transfer entities.Transfer) error
}
