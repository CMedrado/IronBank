package authentication

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	account domain.AccountUseCase
	login   domain.LoginUseCase
}

func NewHandler(accountUseCase domain.AccountUseCase, useCase domain.LoginUseCase) *Handler {
	return &Handler{account: accountUseCase, login: useCase}
}
