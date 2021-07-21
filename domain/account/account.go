package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type UseCase struct {
	StoredAccount Repository
}

//CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func (auc *UseCase) CreateAccount(name string, cpf string, secret string, balance int) (uuid.UUID, error) {
	err, cpf := domain.CheckCPF(cpf)
	if err != nil {
		return uuid.UUID{}, err
	}
	account, err := auc.GetAccountCPF(cpf)
	if err != nil {
		return uuid.UUID{}, err
	}
	err = CheckAccountExistence(account)
	if err != nil {
		return uuid.UUID{}, err
	}
	err = CheckBalance(balance)
	if err != nil {
		return uuid.UUID{}, domain.ErrBalanceAbsent
	}
	aux := 0
	var id uuid.UUID
	for aux == 0 {
		id, _ = domain.Random()
		searchAccount, err := auc.SearchAccount(id)
		if err != nil {
			return uuid.UUID{}, err
		}
		if (searchAccount != domain.Account{}) {
			aux = 0
		} else {
			aux = 1
		}
	}
	secretHash := domain.CreateHash(secret)
	newAccount := domain.Account{ID: id, Name: name, CPF: cpf, Secret: secretHash, Balance: balance, CreatedAt: domain.CreatedAt()}
	err = auc.StoredAccount.SaveAccount(ChangeAccountDomain(newAccount))
	if err != nil {
		return uuid.UUID{}, domain.ErrInsert
	}
	return id, err
}

//GetBalance requests the salary for the Story by sending the ID
func (auc *UseCase) GetBalance(id string) (int, error) {
	idUUID, err := uuid.Parse(id)

	if err != nil {
		return 0, domain.ErrParse
	}

	account, err := auc.SearchAccount(idUUID)
	if err != nil {
		return 0, err
	}
	err = domain.CheckExistID(account)

	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

//GetAccounts returns all API accounts
func (auc *UseCase) GetAccounts() ([]domain.Account, error) {
	accounts, err := auc.StoredAccount.ReturnAccounts()

	if err != nil {
		return []domain.Account{}, domain.ErrInsert
	}

	var account []domain.Account

	for _, a := range accounts {
		account = append(account, ChangeAccountStorage(a))
	}

	return account, nil
}

// SearchAccount returns the account via the received ID
func (auc UseCase) SearchAccount(id uuid.UUID) (domain.Account, error) {
	accounts, err := auc.StoredAccount.ReturnAccounts()
	if err != nil {
		return domain.Account{}, domain.ErrInsert
	}
	account := domain.Account{}

	for _, a := range accounts {
		if a.ID == id {
			account = ChangeAccountStorage(a)
		}
	}

	return account, nil
}

func (auc UseCase) GetAccountCPF(cpf string) (domain.Account, error) {
	accounts, err := auc.StoredAccount.ReturnAccounts()
	if err != nil {
		return domain.Account{}, domain.ErrInsert
	}
	account := domain.Account{}

	for _, a := range accounts {
		if a.CPF == cpf {
			account = ChangeAccountStorage(a)
		}
	}

	return account, nil
}

func (auc UseCase) UpdateBalance(accountOrigin domain.Account, accountDestination domain.Account) error {
	err := auc.StoredAccount.ChangeBalance(ChangeAccountDomain(accountOrigin), ChangeAccountDomain(accountDestination))
	if err != nil {
		return domain.ErrUpdate
	}
	return nil
}
