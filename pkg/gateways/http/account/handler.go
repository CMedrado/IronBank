package account

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	account domain.AccountUseCase
}

func NewHandler(useCase domain.AccountUseCase) *Handler {
	return &Handler{account: useCase}
}
