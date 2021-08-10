package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
)

func (a *Storage) ReturnAccounts() ([]entries.Account, error) {
	statement := `SELECT * FROM accounts`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []entries.Account{}, err
	}
	defer rows.Close()
	var account entries.Account
	var accounts []entries.Account
	for rows.Next() {
		rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
