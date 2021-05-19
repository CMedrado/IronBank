package main

import (
	"github.com/CMedrado/DesafioStone/domain"
	https "github.com/CMedrado/DesafioStone/https"
	"github.com/CMedrado/DesafioStone/store"
	"log"
	"net/http"
)

func main() {
	accountTransfer := store.NewStoredTransferID()
	accountToken := store.NewStoredToked()
	accountLogin := store.NewStoredLogin()
	accountStorage := store.NewStoredAccount()
	accounStorage := domain.AccountUseCase{accountStorage, accountLogin, accountToken, accountTransfer}
	servidor := https.NewServerAccount(&accounStorage)

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
