package accounts

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Storage struct {
	pool *pgxpool.Pool
}

func NewStored(pool *pgxpool.Pool) *Storage {
	return &Storage{pool: pool}
}
