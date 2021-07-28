package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) ReturnAccountCPF(cpf string) (domain.Account, error) {
	var account domain.Account
	statement := `SELECT accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, cpf).Scan(&account)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}
