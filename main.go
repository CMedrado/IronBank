package main

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/authentication"
	"github.com/CMedrado/DesafioStone/domain/transfer"
	"github.com/CMedrado/DesafioStone/https"
	http_account "github.com/CMedrado/DesafioStone/https/account"
	http_login "github.com/CMedrado/DesafioStone/https/authentication"
	http_transfer "github.com/CMedrado/DesafioStone/https/transfer"
	"github.com/CMedrado/DesafioStone/storage/postgre"
	store_account "github.com/CMedrado/DesafioStone/storage/postgre/accounts"
	store_token "github.com/CMedrado/DesafioStone/storage/postgre/token"
	store_transfer "github.com/CMedrado/DesafioStone/storage/postgre/transfer"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const dbFileNameAccount = "accounts.db.json"
const dbFileNameToken = "token.db.json"
const dbFileNameTransfer = "transfer.db.json"
const Port = "5000"
const dbUrl = "postgresql://postgres:example@localhost:5432/desafio?sslmode=disable"

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	lentry := logrus.NewEntry(logger)

	err := postgre.RunMigrations(dbUrl)
	if err != nil {
		lentry.Fatalf("failed to run migrations: %v\n", err)
		os.Exit(1)
	}

	pool, err := getDbPool(dbUrl)
	if err != nil {
		lentry.Fatalf("failed to get db connection pool: %v", err)
	}

	accountStorage := store_account.NewStored(pool, lentry)
	accountToken := store_token.NewStored(pool, lentry)
	accountTransfer := store_transfer.NewStored(pool, lentry)
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

func getDbPool(dburl string) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(dburl)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func openFiles(dbFileNameAccount string, dbFileNameToken string, dbFileNameTransfer string, lentry *logrus.Entry) (*os.File, *os.File, *os.File) {
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

	return dbAccount, dbToken, dbTransfer
}
