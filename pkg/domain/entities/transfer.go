package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	OriginAccountID      uuid.UUID `json:"origin_account_id"`
	DestinationAccountID uuid.UUID `json:"destination_account_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}
