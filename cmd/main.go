package main

import (
	"context"
	account2 "github.com/CMedrado/DesafioStone/pkg/domain/account"
	authentication2 "github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	transfer2 "github.com/CMedrado/DesafioStone/pkg/domain/transfer"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/postgre"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/postgre/accounts"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/postgre/token"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/postgre/transfer"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	"github.com/CMedrado/DesafioStone/pkg/gateways/http/account"
	"github.com/CMedrado/DesafioStone/pkg/gateways/http/authentication"
	transfer3 "github.com/CMedrado/DesafioStone/pkg/gateways/http/transfer"
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
const dbUrl = "postgresql://postgres:example@db:5432/desafio?sslmode=disable"

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	lentry := logrus.NewEntry(logger)

	err := postgre.RunMigrations(dbUrl)
	if err != nil {
		lentry.Fatalf("failed to run migrations: %v\n", err)
	}

	pool, err := getDbPool(dbUrl)
	if err != nil {
		lentry.Fatalf("failed to get db connection pool: %v", err)
	}

	accountStorage := accounts.NewStored(pool, lentry)
	accountToken := token.NewStored(pool, lentry)
	accountTransfer := transfer.NewStored(pool, lentry)
	accountUseCase := account2.UseCase{StoredAccount: accountStorage}
	loginUseCase := authentication2.UseCase{StoredToken: accountToken}
	transferUseCase := transfer2.UseCase{StoredTransfer: accountTransfer}
	accountHandler := account.NewHandler(&accountUseCase, lentry)
	loginHandler := authentication.NewHandler(&accountUseCase, &loginUseCase, lentry)
	transferHandler := transfer3.NewHandler(&accountUseCase, &loginUseCase, &transferUseCase, lentry)
	server := http2.NewAPI(accountHandler, loginHandler, transferHandler, lentry)

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
