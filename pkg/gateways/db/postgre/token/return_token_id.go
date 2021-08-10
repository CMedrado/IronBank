package token

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

func (a *Storage) ReturnTokenID(id uuid.UUID) (entries.Token, error) {
	var token entries.Token
	statement := `SELECT * FROM tokens WHERE id_token=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&token.ID, &token.IdAccount, &token.CreatedAt)
	if err != nil {
		return entries.Token{}, err
	}
	return token, nil
}
