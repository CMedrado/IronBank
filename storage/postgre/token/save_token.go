package token

import (
	"context"
	"github.com/google/uuid"
)

func (a *Storage) SaveToken(idToken uuid.UUID, id uuid.UUID, time string) error {
	statement := `INSERT INTO accounts(id, id_account, cr	seated_at)
				  VALUES ($1, $2, $3)`
	_, err := a.pool.Exec(context.Background(), statement, idToken, id, time)
	if err != nil {
		return err
	}
	return nil
}
