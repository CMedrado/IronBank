package token

import "context"

func (a *Storage) ReturnTokens() ([]Token, error) {
	statement := `SELECT * FROM tokens`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []Token{}, err
	}
	defer rows.Close()
	var token Token
	var tokens []Token
	for rows.Next() {
		rows.Scan(&token.ID, &token.IdAccount, &token.CreatedAt)
		tokens = append(tokens, token)
	}
	return tokens, nil
}
