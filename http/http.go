package http

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/gorilla/mux.v1.8.0"
)

func NewServerAccount(storaged domain.MethodsDomain) *ServerAccount {
	s := new(ServerAccount)

	s.storage = storaged

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{cpf}/balance", s.GetBalance).Methods("GET")
	router.HandleFunc("/accounts", s.GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/", s.CreatedAccount).Methods("POST")
	router.HandleFunc("/login", s.AuthenticatedLogin).Methods("POST")
	router.HandleFunc("/transfers", s.GetTransfers).Methods("GET")
	router.HandleFunc("/transfers", s.MakeTransfers).Methods("POST")

	s.Handler = router

	return s
}
