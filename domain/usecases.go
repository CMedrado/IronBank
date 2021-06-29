package domain

type AccountUseCase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (int, error)
	GetBalance(id int) (int, error)
	GetAccounts() []Account
	SearchAccount(id int) Account
	UpdateBalance(accountOrigin Account, accountDestination Account)
	GetAccountCPF(cpf string) Account
}

type LoginUseCase interface {
	AuthenticatedLogin(cpf, secret string) (error, string)
	GetTokenID(id int) Token
}

type TransferUseCase interface {
	GetTransfers(token string) ([]Transfer, error)
	CreateTransfers(token string, accountDestinationID int, amount int) (error, int)
}
