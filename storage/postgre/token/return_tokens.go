package token

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) ReturnTokens() ([]domain.Token, error) {
	statement := `SELECT * FROM tokens`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []domain.Token{}, err
	}
	defer rows.Close()
	var token domain.Token
	var tokens []domain.Token
	for rows.Next() {
		rows.Scan(&token.ID, &token.IdAccount, &token.CreatedAt)
		tokens = append(tokens, token)
	}
	return tokens, nil
}
