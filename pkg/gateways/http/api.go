package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
)

type API struct {
	account  AccountHandler
	Login    LoginHandler
	transfer TransferHandler
	logger   *zap.Logger
	http.Handler
}

func NewAPI(account AccountHandler, login LoginHandler, transfer TransferHandler, logger *zap.Logger) *API {
	s := new(API)

	s.account = account
	s.Login = login
	s.transfer = transfer
	s.logger = logger

	r := chi.NewRouter()

	r.Use(requestLogger(logger))

	// Directory
	r.Route("/account", func(r chi.Router) {
		r.Post("/", s.account.CreateAccount)
		r.Get("/", s.account.ListAccounts)
		r.Get("/{cpf}", s.account.GetAccount)
		r.Get("/{id}/balance", s.account.GetBalance)
	})

	// Login
	r.Post("/login", s.Login.Login)

	// Transfers
	r.Route("/transfers", func(r chi.Router) {
		r.Post("/", s.transfer.CreateTransfer)
		r.Get("/", s.transfer.ListTransfers)
		r.Get("/transfers/statistic", s.transfer.GetStatisticTransfers)
		r.Get("/transfers/rank", s.transfer.GetRankTransfer)
	})

	s.Handler = r

	return s
}

func requestLogger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(logger.WithCtx(r.Context(), log))
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
