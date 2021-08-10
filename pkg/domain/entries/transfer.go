package entries

import (
	"github.com/google/uuid"
	"time"
)

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	OriginAccountID      uuid.UUID `json:"origin_account_id"`
	DestinationAccountID uuid.UUID `json:"destination_account_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}
