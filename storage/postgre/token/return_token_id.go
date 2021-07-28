package token

import (
	"context"
	"github.com/google/uuid"
)

func (a *Storage) ReturnTokenID(id uuid.UUID) (Token, error) {
	var token Token
	statement := `SELECT tokens WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&token)
	if err != nil {
		return Token{}, err
	}
	return token, nil
}
