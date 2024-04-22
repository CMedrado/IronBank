package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/configuration"
	logger2 "github.com/CMedrado/DesafioStone/pkg/common/logger"
	account2 "github.com/CMedrado/DesafioStone/pkg/domain/account"
	authentication2 "github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	transfer2 "github.com/CMedrado/DesafioStone/pkg/domain/transfer"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/postgre"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/postgre/entries"
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
	logger := logger2.CreateLogger()

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dbURL := config.Database.GetUrl()
	err = postgre.RunMigrations(dbURL)
	if err != nil {
		logger.Error("failed to run migrations", zap.Error(err))
	}

	pool, err := getDbPool(dbURL)
	if err != nil {
		logger.Error("failed to get db connection pool", zap.Error(err))
	}

	accountStorage := entries.NewStored(pool)

	accountUseCase := account2.NewUseCase(accountStorage, client)
	loginUseCase := authentication2.NewUseCase(accountStorage)
	transferUseCase := transfer2.NewUseCase(accountStorage, client)

	accountHandler := account.NewHandler(accountUseCase)
	loginHandler := authentication.NewHandler(loginUseCase)
	transferHandler := transfer3.NewHandler(transferUseCase)

	server := http2.NewAPI(accountHandler, loginHandler, transferHandler, logger)

	serverLog := logger.With(zap.String("Port", config.Api.Port))
	serverLog.Info("starting the server!")
	address := fmt.Sprintf(":%s", config.Api.Port)
	if err = http.ListenAndServe(address, server); err != nil { //nolint:gosec
		logger.Error("could not hear on port", zap.String("Port", config.Api.Port), zap.Error(err))
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
