package authentication

import (
	"github.com/CMedrado/DesafioStone/domain"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	login  domain.LoginUseCase
	logger *log.Entry
}

func NewHandler(useCase domain.LoginUseCase, logger *log.Entry) *Handler {
	return &Handler{login: useCase, logger: logger}
}
