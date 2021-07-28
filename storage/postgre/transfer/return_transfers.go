package transfer

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) ReturnTransfer() ([]domain.Transfer, error) {
	statement := `SELECT * FROM transfers`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []domain.Transfer{}, err
	}
	defer rows.Close()
	var transfer domain.Transfer
	var transfers []domain.Transfer
	for rows.Next() {
		rows.Scan(&transfer.ID, &transfer.OriginAccountID, &transfer.DestinationAccountID, &transfer.Amount, &transfer.CreatedAt)
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
