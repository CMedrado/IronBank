package http

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/gorilla/mux.v1.8.0"
)

func NewServerAccount(storaged domain.StorageMethods) *ServerAccount {
	s := new(ServerAccount)

	s.storaged = storaged

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{cpf}/balance", s.GetBalance).Methods("GET")
	router.HandleFunc("/accounts", s.GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/", s.CreatedAccount).Methods("POST")

	s.Handler = router

	return s
}
