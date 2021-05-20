package https

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/gorilla/mux"
)

func NewServerAccount(storage domain.MethodsDomain) *ServerAccount {
	s := new(ServerAccount)

	s.storage = storage

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}/balance", s.handleBalance).Methods("GET")
	router.HandleFunc("/accounts", s.handleAccounts).Methods("GET")
	router.HandleFunc("/accounts", s.processAccount).Methods("POST")
	router.HandleFunc("/login", s.processLogin).Methods("POST")
	router.HandleFunc("/transfers", s.handleTransfers).Methods("GET")
	router.HandleFunc("/transfers", s.processTransfer).Methods("POST")

	s.Handler = router

	return s
}
