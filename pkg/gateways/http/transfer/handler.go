package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	account  domain.AccountUseCase
	login    domain.LoginUseCase
	transfer domain.TransferUseCase
}

func NewHandler(accountUseCase domain.AccountUseCase, loginUseCase domain.LoginUseCase, useCase domain.TransferUseCase) *Handler {
	return &Handler{account: accountUseCase, login: loginUseCase, transfer: useCase}
}
