package token

import "encoding/json"

func (a *StoredToken) PostToken(id int, token string) {
	a.tokens = append(a.tokens, Token{ID: id, Token: token})
	a.dataBase.Seek(0, 0)

	json.NewEncoder(a.dataBase).Encode(a.tokens)
}

func (a *StoredToken) GetTokens() []Token {
	return a.tokens
}
