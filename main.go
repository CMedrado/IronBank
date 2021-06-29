package main

import (
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/login"
	"github.com/CMedrado/DesafioStone/domain/transfer"
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
	accountStorage := store.NewStoredAccount()
	accountUseCase := account.UseCase{StoredAccount: accountStorage}
	loginUseCase := login.UseCase{AccountUseCase: &accountUseCase, StoredToken: accountToken}
	transferUseCase := transfer.UseCase{AccountUseCase: &accountUseCase, StoredTransfer: accountTransfer, TokenUseCase: &loginUseCase}
	server := https.NewServerAccount(&accountUseCase, &loginUseCase, &transferUseCase, lentry)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
