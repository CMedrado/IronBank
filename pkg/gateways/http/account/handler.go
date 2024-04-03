package account

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	account domain2.AccountUseCase
}

func NewHandler(useCase domain2.AccountUseCase) *Handler {
	return &Handler{account: useCase}
}
