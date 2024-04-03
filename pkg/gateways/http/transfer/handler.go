package transfer

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	account  domain2.AccountUseCase
	login    domain2.LoginUseCase
	transfer domain2.TransferUseCase
}

func NewHandler(accountUseCase domain2.AccountUseCase, loginUseCase domain2.LoginUseCase, useCase domain2.TransferUseCase) *Handler {
	return &Handler{account: accountUseCase, login: loginUseCase, transfer: useCase}
}
