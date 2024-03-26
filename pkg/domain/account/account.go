package account

import "C"
import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	redis2 "github.com/CMedrado/DesafioStone/pkg/gateways/redis"
)

type UseCase struct {
	StoredAccount Repository
	logger        *logrus.Entry
	redis         *redis.Client
}

// CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
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
		return uuid.UUID{}, ErrBalanceAbsent
	}
	id, _ := domain.Random()
	secretHash := domain.CreateHash(secret)
	newAccount := entities.Account{ID: id, Name: name, CPF: cpf, Secret: secretHash, Balance: balance, CreatedAt: domain.CreatedAt()}
	err = auc.StoredAccount.SaveAccount(newAccount)
	if err != nil {
		return uuid.UUID{}, domain.ErrInsert
	}
	return id, err
}

// GetBalance requests the salary for the Story by sending the ID
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

// GetAccounts returns all API accounts
func (auc *UseCase) GetAccounts() ([]entities.Account, error) {
	accounts, err := auc.StoredAccount.ReturnAccounts()

	if err != nil {
		return []entities.Account{}, domain.ErrInsert
	}

	return accounts, nil
}

// SearchAccount returns the account via the received ID
func (auc UseCase) SearchAccount(id uuid.UUID) (entities.Account, error) {
	account, err := auc.StoredAccount.ReturnAccountID(id)
	if err != nil {
		return entities.Account{}, domain.ErrSelect
	}
	return account, nil
}

func (auc UseCase) GetAccountCPF(cpf string) (entities.Account, error) {
	account, err := redis2.Get(context.Background(), cpf, auc.redis)

	if err != nil && !errors.Is(err, domain.ErrAccountNotFound) {
		return entities.Account{}, err
	}

	if errors.Is(err, domain.ErrAccountNotFound) {
		auc.logger.Print("Storing")
		account, err = auc.StoredAccount.ReturnAccountCPF(cpf)
		if err != nil {
			return entities.Account{}, domain.ErrSelect
		}
		err = redis2.Set(context.Background(), account, auc.redis)
		if err != nil {
			return entities.Account{}, err
		}
	}

	return account, nil
}

func (auc UseCase) UpdateBalance(accountOrigin entities.Account, accountDestination entities.Account) error {
	err := auc.StoredAccount.ChangeBalance(accountOrigin, accountDestination)
	if err != nil {
		return ErrUpdate
	}
	return nil
}

func NewUseCase(repository Repository, log *logrus.Entry, redis *redis.Client) *UseCase {
	return &UseCase{StoredAccount: repository, logger: log, redis: redis}
}
