package store

type Account struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Secret    string `json:"secret"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type StoredAccount struct {
	storage map[string]Account
}
