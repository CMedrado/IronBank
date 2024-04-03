package authentication

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	account domain2.AccountUseCase
	login   domain2.LoginUseCase
}

func NewHandler(accountUseCase domain2.AccountUseCase, useCase domain2.LoginUseCase) *Handler {
	return &Handler{account: accountUseCase, login: useCase}
}
