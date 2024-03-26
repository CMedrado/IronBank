package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"github.com/CMedrado/DesafioStone/pkg/common/configuration"
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
)

func main() {
	config, err := configuration.LoadConfigs()
	if err != nil {
		fmt.Printf("something went wrong when reading config env vars: %v\n", err)
		os.Exit(1)
	}
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	logLevel, err := logrus.ParseLevel(config.Api.LogLevel)
	if err != nil {
		fmt.Printf("informed log level on config is invalid: %v", err)
		os.Exit(1)
	}
	logger.SetLevel(logLevel)
	lentry := logrus.NewEntry(logger)

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dbURL := config.Database.GetUrl()
	err = postgre.RunMigrations(dbURL)
	if err != nil {
		lentry.Fatalf("failed to run migrations: %v\n", err)
	}

	pool, err := getDbPool(dbURL)
	if err != nil {
		lentry.Fatalf("failed to get db connection pool: %v", err)
	}

	accountStorage := accounts.NewStored(pool, lentry)
	accountToken := token.NewStored(pool, lentry)
	accountTransfer := transfer.NewStored(pool, lentry)
	accountUseCase := account2.NewUseCase(accountStorage, lentry, client)
	loginUseCase := authentication2.NewUseCase(accountToken, lentry)
	transferUseCase := transfer2.NewUseCase(accountTransfer, lentry)
	accountHandler := account.NewHandler(accountUseCase, lentry)
	loginHandler := authentication.NewHandler(accountUseCase, loginUseCase, lentry)
	transferHandler := transfer3.NewHandler(accountUseCase, loginUseCase, transferUseCase, lentry)
	server := http2.NewAPI(accountHandler, loginHandler, transferHandler, lentry)

	serverLog := lentry.WithField("Port", config.Api.Port)
	serverLog.Info("starting the server!")
	address := fmt.Sprintf(":%s", config.Api.Port)
	if err = http.ListenAndServe(address, server); err != nil { //nolint:gosec
		lentry.Fatalf("could not hear on port %s: %v", config.Api.Port, err)
	}
	serverLog.Info("shutting down the server")
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

//func openFiles(dbFileNameAccount string, dbFileNameToken string, dbFileNameTransfer string, lentry *logrus.Entry) (*os.File, *os.File, *os.File) {
//	dbAccount, err := os.OpenFile(dbFileNameAccount, os.O_RDWR|os.O_CREATE, 0666)
//
//	if err != nil {
//		lentry.Fatalf("problem opening %s %v", dbFileNameAccount, err)
//	}
//
//	dbToken, err := os.OpenFile(dbFileNameToken, os.O_RDWR|os.O_CREATE, 0666)
//
//	if err != nil {
//		lentry.Fatalf("problem opening %s %v", dbFileNameToken, err)
//	}
//
//	dbTransfer, err := os.OpenFile(dbFileNameTransfer, os.O_RDWR|os.O_CREATE, 0666)
//
//	if err != nil {
//		lentry.Fatalf("problem opening %s %v", dbFileNameTransfer, err)
//	}
//
//	return dbAccount, dbToken, dbTransfer
//}
