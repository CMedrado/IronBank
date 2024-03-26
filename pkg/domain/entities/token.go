package entities

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID `json:"id"`
	IdAccount uuid.UUID `json:"id_account"`
	CreatedAt time.Time `json:"created_at"`
}
