package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Token struct {
	ID        uuid.UUID `json:"id"`
	IdAccount uuid.UUID `json:"id_account"`
	CreatedAt time.Time `json:"created_at"`
}

type Storage struct {
	pool *pgxpool.Pool
}

func NewStored(pool *pgxpool.Pool) *Storage {
	return &Storage{pool: pool}
}
