package http

import (
	"github.com/CMedrado/DesafioStone/domain"
	"net/http"
)

type CreatedRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenRequest struct {
	Token           int `json:"token"`
	AccountOriginID int `json:"account_origin_id"`
}

type TransfersRequest struct {
	Token                int `json:"token"`
	AccountOriginID      int `json:"account_origin_id"`
	AccountDestinationID int `json:"account_destination_id"`
	Amount               int `json:"amount"`
	ID                   int `json:"id"`
}

type ServerAccount struct {
	storage domain.MethodsDomain
	http.Handler
}
