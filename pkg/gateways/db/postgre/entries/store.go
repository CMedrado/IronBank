package entries

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

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	OriginAccountID      uuid.UUID `json:"origin_account_id"`
	DestinationAccountID uuid.UUID `json:"destination_account_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

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
