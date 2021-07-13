package token

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Token struct {
	ID        int       `json:"id"`
	IdAccount uuid.UUID `json:"id_account"`
	CreatedAt string    `json:"created_at"`
}

type Storage struct {
	pool *pgxpool.Pool

	log *logrus.Entry
}

func NewStored(pool *pgxpool.Pool, log *logrus.Entry) *Storage {
	return &Storage{pool: pool, log: log}
}
