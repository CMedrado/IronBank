package store

var accountStorage = make(map[string]Account)
var accountLogin = make(map[int]Login)
var accountTransferAccountID = make(map[int]StoredTransferID)
var accountToken = make(map[int]Token)
var accountTransferID = make(map[int]Transfer)

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

type Token struct {
	Token string `json:"token"`
}

type Transfer struct {
	ID                   int    `json:"id"`
	AccountOriginID      int    `json:"account_origin_id"`
	AccountDestinationID int    `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}

type StoredTransferID struct {
	accountTransferID map[int]Transfer
}

type StoredTransferAccountID struct {
	accountTransferAccountID map[int]StoredTransferID
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

func NewStoredTransferAccountID() *StoredTransferAccountID {
	return &StoredTransferAccountID{accountTransferAccountID}
}

func NewStoredToked() *StoredToken {
	return &StoredToken{accountToken}
}

func NewStoredTransferID() *StoredTransferID {
	return &StoredTransferID{accountTransferID}
}
