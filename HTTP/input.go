package HTTP

import (
	"github.com/CMedrado/DesafioStone/Domain"
	"net/http"
)

type CreatedRequest struct {
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Secret    string `json:"secret"`
}

type ServidorConta struct {
	armazenamento Domain.MetodosDeArmazenamento
	http.Handler
}
