package http

import "net/http"

type AccountHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	ListAccounts(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
}

type LoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type TransferHandler interface {
	ListTransfers(w http.ResponseWriter, r *http.Request)
	CreateTransfer(w http.ResponseWriter, r *http.Request)
	GetCountTransfer(w http.ResponseWriter, r *http.Request)
	GetRankTransfer(w http.ResponseWriter, r *http.Request)
}
