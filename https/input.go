package https

import (
	"github.com/CMedrado/DesafioStone/domain"
	"net/http"
)

type CreateRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance uint   `json:"balance"`
}

type CreateResponse struct {
	ID int `json:"id"`
}

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TransfersRequest struct {
	AccountOriginID      int  `json:"account_origin_id"`
	Token                int  `json:"token"`
	AccountDestinationID int  `json:"account_destination_id"`
	Amount               uint `json:"amount"`
}

type TransferResponse struct {
	ID int `json:"id"`
}

type ServerAccount struct {
	storage domain.MethodsDomain
	http.Handler
}
