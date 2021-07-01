package main

import (
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/authentication"
	"github.com/CMedrado/DesafioStone/domain/transfer"
	https "github.com/CMedrado/DesafioStone/https"
	http_account "github.com/CMedrado/DesafioStone/https/account"
	http_login "github.com/CMedrado/DesafioStone/https/authentication"
	http_transfer "github.com/CMedrado/DesafioStone/https/transfer"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	store_token "github.com/CMedrado/DesafioStone/storage/file/token"
	store_transfer "github.com/CMedrado/DesafioStone/storage/file/transfer"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	time "time"
)

const dbFileNameAccount = "accounts.db.json"
const dbFileNameToken = "token.db.json"
const dbFileNameTransfer = "transfer.db.json"
const Port = "5000"

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	lentry := logrus.NewEntry(logger)

	dbAccount, err := os.OpenFile(dbFileNameAccount, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		lentry.Fatal("problem opening %s %v", dbFileNameAccount, err)
	}

	dbToken, err := os.OpenFile(dbFileNameToken, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		lentry.Fatal("problem opening %s %v", dbFileNameToken, err)
	}

	dbTransfer, err := os.OpenFile(dbFileNameTransfer, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		lentry.Fatal("problem opening %s %v", dbFileNameTransfer, err)
	}

	accountStorage := store_account.NewStoredAccount(dbAccount)
	accountToken := store_token.NewStoredToked(dbToken)
	accountTransfer := store_transfer.NewStoredTransfer(dbTransfer)
	accountUseCase := account.UseCase{StoredAccount: accountStorage}
	loginUseCase := authentication.UseCase{AccountUseCase: &accountUseCase, StoredToken: accountToken}
	transferUseCase := transfer.UseCase{AccountUseCase: &accountUseCase, StoredTransfer: accountTransfer, TokenUseCase: &loginUseCase}
	accountHandler := http_account.NewHandler(&accountUseCase, lentry)
	loginHandler := http_login.NewHandler(&loginUseCase, lentry)
	transferHandler := http_transfer.NewHandler(&transferUseCase, lentry)
	server := https.NewAPI(accountHandler, loginHandler, transferHandler, lentry)
	lentry.WithField("Port", Port).Info("starting the server!")
	if err := http.ListenAndServe(":5000", server); err != nil {
		lentry.Fatal("could not hear on port 5000 ")
	}
	lentry.WithField("Port", Port).Info("shutting down the server")
}
