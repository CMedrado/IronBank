package store

var accountStorage = make(map[string]Account)
var accountLogin = make(map[int]Login)
var accountTransfer = make(map[int]map[int]Transfer)

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
	Token  int    `json:"token"`
}

type Transfer struct {
	ID                   int    `json:"id"`
	AccountOriginID      int    `json:"accountoriginid"`
	AccountDestinationID string `json:"accountdestinationid"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}

type StoredTransfer struct {
	accountTransfer map[int]map[int]Transfer
}

type StoredLogin struct {
	accountLogin map[int]Login
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

func NewStoredTransfer() *StoredTransfer {
	return &StoredTransfer{accountTransfer}
}
