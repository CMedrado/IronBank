package transfer

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	account  domain2.AccountUseCase
	login    domain2.LoginUseCase
	transfer domain2.TransferUseCase
	logger   *logrus.Entry
}

func NewHandler(accountUseCase domain2.AccountUseCase, loginUseCase domain2.LoginUseCase, useCase domain2.TransferUseCase, logger *logrus.Entry) *Handler {
	return &Handler{account: accountUseCase, login: loginUseCase, transfer: useCase, logger: logger}
}
