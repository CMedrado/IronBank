package token

import (
	"context"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func (a *Storage) SaveToken(token entities.Token) error {
	statement := `INSERT INTO tokens(id_token, id_account, created_at)
				  VALUES ($1, $2, $3)`
	comand, err := a.pool.Exec(context.Background(), statement, token.ID, token.IdAccount, token.CreatedAt)
	if comand.RowsAffected() > 0 {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
