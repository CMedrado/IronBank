package domain

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

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

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
