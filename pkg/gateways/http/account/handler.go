package account

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	account domain2.AccountUseCase
	logger  *log.Entry
}

func NewHandler(useCase domain2.AccountUseCase, logger *log.Entry) *Handler {
	return &Handler{account: useCase, logger: logger}
}
