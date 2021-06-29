package https

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CreateRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TransfersRequest struct {
	AccountOriginID      string `json:"account_origin_id"`
	Token                int    `json:"token"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}

type TransferResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetTransfersResponse struct {
	Transfers []domain.Transfer `json:"transfers"`
}

type GetAccountsResponse struct {
	Accounts []domain.Account `json:"accounts"`
}

type ServerAccount struct {
	account  domain.AccountUseCase
	login    domain.LoginUseCase
	transfer domain.TransferUseCase
	logger   *logrus.Entry

	http.Handler
}

type ErrorsResponse struct {
	Errors string `json:"errors"`
}
