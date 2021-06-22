package main

import (
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/login"
	"github.com/CMedrado/DesafioStone/domain/transfer"
	https "github.com/CMedrado/DesafioStone/https"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	store_token "github.com/CMedrado/DesafioStone/storage/file/token"
	store_transfer "github.com/CMedrado/DesafioStone/storage/file/transfer"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	time "time"
)

const dbFileNameAccount = "accounts.db.json"
const dbFileNameToken = "token.db.json"
const dbFileNameTransfer = "transfer.db.json"

func main() {
	dbAccount, err := os.OpenFile(dbFileNameAccount, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal("problem opening %s %v", dbFileNameAccount, err)
	}

	dbToken, err := os.OpenFile(dbFileNameToken, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal("problem opening %s %v", dbFileNameToken, err)
	}

	dbTransfer, err := os.OpenFile(dbFileNameTransfer, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal("problem opening %s %v", dbFileNameTransfer, err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	lentry := logrus.NewEntry(logger)

	accountStorage := store_account.NewStoredAccount(dbAccount)
	accountToken := store_token.NewStoredToked(dbToken)
	accountTransfer := store_transfer.NewStoredTransfer(dbTransfer)
	accountUseCase := account.UseCase{StoredAccount: accountStorage}
	loginUseCase := login.UseCase{AccountUseCase: &accountUseCase, StoredToken: accountToken}
	transferUseCase := transfer.UseCase{AccountUseCase: &accountUseCase, StoredTransfer: accountTransfer, TokenUseCase: &loginUseCase}
	server := https.NewServerAccount(&accountUseCase, &loginUseCase, &transferUseCase, lentry)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
