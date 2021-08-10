package http

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	account  AccountHandler
	Login    LoginHandler
	transfer TransferHandler
	logger   *logrus.Entry
	http.Handler
}

func NewAPI(account AccountHandler, login LoginHandler, transfer TransferHandler, logger *logrus.Entry) *API {
	s := new(API)

	s.account = account
	s.Login = login
	s.transfer = transfer
	s.logger = logger

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}/balance", s.account.GetBalance).Methods("GET")
	router.HandleFunc("/accounts", s.account.ListAccounts).Methods("GET")
	router.HandleFunc("/accounts", s.account.CreateAccount).Methods("POST")
	router.HandleFunc("/login", s.Login.Login).Methods("POST")
	router.HandleFunc("/transfers", s.transfer.ListTransfers).Methods("GET")
	router.HandleFunc("/transfers", s.transfer.CreateTransfer).Methods("POST")

	s.Handler = router

	return s
}
