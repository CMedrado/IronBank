package transfer

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	OriginAccountID      uuid.UUID `json:"origin_account_id"`
	DestinationAccountID uuid.UUID `json:"destination_account_id"`
	Amount               int       `json:"amount"`
	CreatedAt            string    `json:"created_at"`
}

type Storage struct {
	pool *pgxpool.Pool

	log *logrus.Entry
}

func NewStored(pool *pgxpool.Pool, log *logrus.Entry) *Storage {
	return &Storage{pool: pool, log: log}
}
