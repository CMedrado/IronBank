package http

import (
	"github.com/CMedrado/DesafioStone/domain"
	"net/http"
)

type CreatedRequest struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type ServerAccount struct {
	storaged domain.MethodsDomain
	http.Handler
}
