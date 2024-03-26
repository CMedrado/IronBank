package accounts

import (
	"context"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func (a *Storage) ChangeBalance(accountOrigin, accountDestination entities.Account) error {
	aO := ChangeAccountDomain(accountOrigin)
	aD := ChangeAccountDomain(accountDestination)
	statement := `UPDATE accounts
				  SET balance=$1
				  WHERE name=$2`
	_, err := a.pool.Exec(context.Background(), statement, aO.Balance, aO.Name)
	if err != nil {
		return err
	}
	_, err = a.pool.Exec(context.Background(), statement, aD.Balance, aD.Name)
	if err != nil {
		return err
	}
	return nil
}
