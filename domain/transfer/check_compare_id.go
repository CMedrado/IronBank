package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

// CheckCompareID Compare two IDs to see if they are the same and returns nil if not, it returns an error
func CheckCompareID(accountOriginID, accountDestinationID uuid.UUID) error {
	if accountOriginID == accountDestinationID {
		return domain.ErrSameAccount
	}
	return nil
}
