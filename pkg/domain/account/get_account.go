package account

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	redis2 "github.com/CMedrado/DesafioStone/pkg/gateways/redis"
)

// GetAccounts returns all API accounts
func (auc UseCase) GetAccounts() ([]entities.Account, error) {
	accounts, err := auc.StoredAccount.ReturnAccounts()
	if err != nil {
		return []entities.Account{}, domain.ErrSelect
	}

	return accounts, nil
}

func (auc UseCase) GetAccountCPF(ctx context.Context, cpf string) (entities.Account, error) {
	err, cpf := domain.CheckCPF(cpf)
	if err != nil {
		return entities.Account{}, err
	}

	account, err := redis2.Get(ctx, cpf, auc.redis)
	if err != nil && !errors.Is(err, domain.ErrAccountNotFound) {
		return entities.Account{}, err
	}

	if errors.Is(err, domain.ErrAccountNotFound) {
		account, err = auc.StoredAccount.ReturnAccountCPF(cpf)
		if err != nil {
			return entities.Account{}, domain.ErrSelect
		}

		err = redis2.Set(ctx, account, auc.redis)
		if err != nil {
			return entities.Account{}, err
		}
	}

	return account, nil
}

func (auc UseCase) GetAccountID(id uuid.UUID) (entities.Account, error) {
	account, err := auc.StoredAccount.ReturnAccountID(id)
	if err != nil {
		return entities.Account{}, domain.ErrSelect
	}
	return account, nil
}
