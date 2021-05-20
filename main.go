package main

import (
	"github.com/CMedrado/DesafioStone/domain"
	https "github.com/CMedrado/DesafioStone/https"
	"github.com/CMedrado/DesafioStone/store"
	"log"
	"net/http"
)

func main() {
	accountTransfer := store.NewStoredTransferAccountID()
	accountToken := store.NewStoredToked()
	accountLogin := store.NewStoredLogin()
	accountStorage := store.NewStoredAccount()
	accountUseCase := domain.AccountUseCase{Store: accountStorage, Login: accountLogin, Token: accountToken, Transfer: accountTransfer}
	server := https.NewServerAccount(&accountUseCase)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
