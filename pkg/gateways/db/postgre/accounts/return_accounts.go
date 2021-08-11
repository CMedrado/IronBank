package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func (a *Storage) ReturnAccounts() ([]entities.Account, error) {
	statement := `SELECT * FROM accounts`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []entities.Account{}, err
	}
	defer rows.Close()
	var account entities.Account
	var accounts []entities.Account
	for rows.Next() {
		rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
