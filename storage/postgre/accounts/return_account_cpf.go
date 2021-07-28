package accounts

import "context"

func (a *Storage) ReturnAccountCPF(cpf string) (Account, error) {
	var account Account
	statement := `SELECT accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, cpf).Scan(&account)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}
