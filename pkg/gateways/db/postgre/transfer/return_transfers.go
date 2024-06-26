package transfer

import (
	"context"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func (a *Storage) ReturnTransfer(id uuid.UUID) ([]entities.Transfer, error) {
	statement := `SELECT * FROM transfers WHERE origin_account_id=$1`
	rows, err := a.pool.Query(context.Background(), statement, id)
	if err != nil && err.Error() != ("no rows in result set") {
		return []entities.Transfer{}, err
	}
	defer rows.Close()
	var transfer entities.Transfer
	var transfers []entities.Transfer
	for rows.Next() {
		err = rows.Scan(&transfer.ID, &transfer.OriginAccountID, &transfer.DestinationAccountID, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
