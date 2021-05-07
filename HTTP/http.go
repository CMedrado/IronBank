package HTTP

import (
	"github.com/CMedrado/DesafioStone/Domain"
	"net/http"
)

func NovoServidorConta(armazenamento Domain.MetodosDeArmazenamento) *ServidorConta{
	s := new(ServidorConta)

	s.armazenamento = armazenamento

	roteador := http.NewServeMux()
	//roteador.Handle("/accounts/{id}/balance", http.HandlerFunc(s.AçãoSaldo))
	roteador.Handle("/accounts", http.HandlerFunc(s.AçãoMostrarContas))
	roteador.Handle("/accounts/", http.HandlerFunc(s.AçãoCriarConta))

	s.Handler = roteador

	return s
}

//func (s *ServidorConta) AçãoSaldo(w http.ResponseWriter, r *http.Request) {
	//balance := s.armazenamento.MostrarSaldo(conta)
	//fmt.Fprint(w, balance)
//}

func (s *ServidorConta) AçãoMostrarContas(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	//json.NewDecoder(w).Encode(s.())
}

