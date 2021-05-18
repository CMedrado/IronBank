package https

import (
	"github.com/CMedrado/DesafioStone/domain"
	"net/http"
)

type CreatedRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance uint   `json:"balance"`
}

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenRequest struct {
	Token           string `json:"token"`
	AccountOriginID int    `json:"accountoriginid"`
}

type TransfersRequest struct {
	AccountOriginID      int  `json:"accountoriginid"`
	Token                int  `json:"token"`
	AccountDestinationID int  `json:"accountdestinationid"`
	Amount               uint `json:"amount"`
	ID                   int  `json:"id"`
}

type ServerAccount struct {
	storage domain.MethodsDomain
	http.Handler
}
