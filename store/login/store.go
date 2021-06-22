package login

var accountLogin = make(map[int]Login)

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type StoredLogin struct {
	accountLogin map[int]Login
}

func NewStoredLogin() *StoredLogin {
	return &StoredLogin{accountLogin: accountLogin}
}
