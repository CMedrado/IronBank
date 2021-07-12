package accounts

import "context"

func (a *Storage) ChangeBalance(person1, person2 Account) error {
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
