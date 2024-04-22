package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	transfer domain.TransferUseCase
}

func NewHandler(useCase domain.TransferUseCase) *Handler {
	return &Handler{transfer: useCase}
}
