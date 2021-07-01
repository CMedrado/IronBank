package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	transfer domain.TransferUseCase
	logger   *logrus.Entry
}

func NewHandler(useCase domain.TransferUseCase, logger *logrus.Entry) *Handler {
	return &Handler{transfer: useCase, logger: logger}
}
