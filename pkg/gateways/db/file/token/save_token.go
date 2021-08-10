package token

import (
	"encoding/json"
	"github.com/google/uuid"
)

func (a *StoredToken) SaveToken(id uuid.UUID, token string) {
	a.tokens = append(a.tokens, Token{ID: id, Token: token})
	a.dataBase.Seek(0, 0)

	json.NewEncoder(a.dataBase).Encode(a.tokens)
}
