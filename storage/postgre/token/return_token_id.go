package token

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

func (a *Storage) ReturnTokenID(id uuid.UUID) (domain.Token, error) {
	var token domain.Token
	statement := `SELECT tokens WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&token)
	if err != nil {
		return domain.Token{}, err
	}
	return token, nil
}
