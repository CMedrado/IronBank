package authentication

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	account domain2.AccountUseCase
	login   domain2.LoginUseCase
	logger  *log.Entry
}

func NewHandler(accountUseCase domain2.AccountUseCase, useCase domain2.LoginUseCase, logger *log.Entry) *Handler {
	return &Handler{account: accountUseCase, login: useCase, logger: logger}
}
