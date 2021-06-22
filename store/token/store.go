package token

import (
	"io"
)

type Tokens []Token

type Token struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type StoredToken struct {
	dataBase io.ReadWriteSeeker
	tokens   Tokens
}

func NewStoredToked(dataBase io.ReadWriteSeeker) *StoredToken {
	dataBase.Seek(0, 0)
	token, _ := NewToken(dataBase)

	return &StoredToken{dataBase: dataBase, tokens: token}
}
