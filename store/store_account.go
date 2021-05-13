package store

var accountStorage = make(map[string]Account)
var accountLogin = make(map[int]Login)
var accountTransferTwo = make(map[int]map[int]Transfer)
var accountToken = make(map[int]Token)

type Account struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	CPF       string   `json:"cpf"`
	Secret    [32]byte `json:"secret"`
	Balance   int      `json:"balance"`
	CreatedAt string   `json:"created_at"`
}
type Login struct {
	CPF    string   `json:"cpf"`
	Secret [32]byte `json:"secret"`
}

type Token struct {
	Token           int `json:"token"`
	AccountOriginID int `json:"accountoriginid"`
}

type Transfer struct {
	ID                   int    `json:"id"`
	AccountOriginID      int    `json:"accountoriginid"`
	AccountDestinationID int    `json:"accountdestinationid"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}

type StoredTransferTwo struct {
	accountTransferTwo map[int]map[int]Transfer
}

type StoredLogin struct {
	accountLogin map[int]Login
}

type StoredToken struct {
	accountToken map[int]Token
}

type StoredAccount struct {
	accountStorage map[string]Account
}

func NewStoredAccount() *StoredAccount {
	return &StoredAccount{accountStorage}
}

func NewStoredLogin() *StoredLogin {
	return &StoredLogin{accountLogin}
}

func NewStoredTransferTwo() *StoredTransferTwo {
	return &StoredTransferTwo{accountTransferTwo}
}

func NewStoredToked() *StoredToken {
	return &StoredToken{accountToken}
}
