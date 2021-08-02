package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) ReturnAccountCPF(cpf string) (domain.Account, error) {
	var account domain.Account
	statement := `SELECT * FROM accounts WHERE cpf=$1`
	err := a.pool.QueryRow(context.Background(), statement, cpf).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}
