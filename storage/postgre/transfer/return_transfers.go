package transfer

import "context"

func (a *Storage) ReturnTransfer() ([]Transfer, error) {
	statement := `SELECT * FROM transfers`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []Transfer{}, err
	}
	defer rows.Close()
	var transfer Transfer
	var transfers []Transfer
	for rows.Next() {
		rows.Scan(&transfer.ID, &transfer.OriginAccountID, &transfer.DestinationAccountID, &transfer.Amount, &transfer.CreatedAt)
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
