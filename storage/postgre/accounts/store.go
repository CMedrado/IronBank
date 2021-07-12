package accounts

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt string    `json:"created_at"`
}

type Storage struct {
	pool *pgxpool.Pool

	log *logrus.Entry
}

func NewStored(pool *pgxpool.Pool, log *logrus.Entry) *Storage {
	return &Storage{pool: pool, log: log}
}
