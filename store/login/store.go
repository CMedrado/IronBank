package login

var accountLogin = make(map[int]Login)
var accountToken = make(map[int]Token)

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Token struct {
	Token string `json:"token"`
}

type StoredLogin struct {
	accountLogin map[int]Login
}

type StoredToken struct {
	accountToken map[int]Token
}

func NewStoredLogin() *StoredLogin {
	return &StoredLogin{accountLogin}
}

func NewStoredToked() *StoredToken {
	return &StoredToken{accountToken}
}
