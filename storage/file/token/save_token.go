package token

import (
	"encoding/json"
)

func (a *StoredToken) SaveToken(id int, token string) {
	a.tokens = append(a.tokens, Token{ID: id, Token: token})
	a.dataBase.Seek(0, 0)

	json.NewEncoder(a.dataBase).Encode(a.tokens)
}
