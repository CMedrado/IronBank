package token

var accountToken = make(map[int]Token)

type Token struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type StoredToken struct {
	accountToken map[int]Token
}

func NewStoredToked() *StoredToken {
	return &StoredToken{accountToken}
}
