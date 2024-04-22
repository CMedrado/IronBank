package authentication

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
)

type Handler struct {
	login domain.LoginUseCase
}

func NewHandler(useCase domain.LoginUseCase) *Handler {
	return &Handler{login: useCase}
}
