package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
)

func (a *Storage) ReturnAccountCPF(cpf string) (entries.Account, error) {
	var account entries.Account
	statement := `SELECT * FROM accounts WHERE cpf=$1`
	err := a.pool.QueryRow(context.Background(), statement, cpf).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil && err.Error() != ("no rows in result set") {
		return entries.Account{}, err
	}
	return account, nil
}
