package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
)

func (a *Storage) ChangeBalance(personDomain1, personDomain2 entries.Account) error {
	person1 := ChangeAccountDomain(personDomain1)
	person2 := ChangeAccountDomain(personDomain2)
	statement := `UPDATE accounts
				  SET balance=$1
				  WHERE name=$2`
	_, err := a.pool.Exec(context.Background(), statement, person1.Balance, person1.Name)
	if err != nil {
		return err
	}
	a.pool.Exec(context.Background(), statement, person2.Balance, person2.Name)
	if err != nil {
		return err
	}
	return nil
}
