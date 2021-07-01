package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	account domain.AccountUseCase
	logger  *log.Entry
}

func NewHandler(useCase domain.AccountUseCase, logger *log.Entry) *Handler {
	return &Handler{account: useCase, logger: logger}
}
