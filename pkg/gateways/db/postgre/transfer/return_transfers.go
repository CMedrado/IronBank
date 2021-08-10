package transfer

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

func (a *Storage) ReturnTransfer(id uuid.UUID) ([]entries.Transfer, error) {
	statement := `SELECT * FROM transfers WHERE origin_account_id=$1`
	rows, err := a.pool.Query(context.Background(), statement, id)
	if err != nil {
		return []entries.Transfer{}, err
	}
	defer rows.Close()
	var transfer entries.Transfer
	var transfers []entries.Transfer
	for rows.Next() {
		rows.Scan(&transfer.ID, &transfer.OriginAccountID, &transfer.DestinationAccountID, &transfer.Amount, &transfer.CreatedAt)
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
