package main

import (
	"github.com/CMedrado/DesafioStone/domain"
	https "github.com/CMedrado/DesafioStone/https"
	"github.com/CMedrado/DesafioStone/store"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	time "time"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	lentry := logrus.NewEntry(logger)

	accountTransfer := store.NewStoredTransferAccountID()
	accountToken := store.NewStoredToked()
	accountLogin := store.NewStoredLogin()
	accountStorage := store.NewStoredAccount()
	accountUseCase := account.UseCase{Store: accountStorage, Login: accountLogin, Token: accountToken, Transfer: accountTransfer}
	server := https.NewServerAccount(&accountUseCase.Store, &accountUseCase.Login, &accountUseCase.Transfer, lentry)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
