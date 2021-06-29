package domain

import "github.com/google/uuid"

type Account struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt string    `json:"created_at"`
}

type Token struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	AccountOriginID      uuid.UUID `json:"account_origin_id"`
	AccountDestinationID uuid.UUID `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            string    `json:"created_at"`
}

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
