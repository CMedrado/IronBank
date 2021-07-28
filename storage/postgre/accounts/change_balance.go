package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) ChangeBalance(personDomain1, personDomain2 domain.Account) error {
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
