package token

import (
	"encoding/json"
	"io"
)

func NewToken(rdr io.Reader) ([]Token, error) {
	var tokens []Token
	err := json.NewDecoder(rdr).Decode(&tokens)
	if err != nil {
		return tokens, err
	}
	return tokens, err
}
