package entities

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
