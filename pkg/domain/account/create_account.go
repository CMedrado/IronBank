package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/CMedrado/DesafioStone/pkg/common/logger"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

// CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func (auc UseCase) CreateAccount(ctx context.Context, name string, cpf string, secret string, balance int) (uuid.UUID, error) {
	l := logger.FromCtx(ctx).With(
		zap.String("module", "handler"),
		zap.String("method", "createAccount"),
	)

	err, cpf := domain.CheckCPF(cpf)
	if err != nil {
		return uuid.UUID{}, err
	}

	account, err := auc.StoredAccount.ReturnAccountCPF(cpf)
	if err != nil && !errors.Is(err, domain.ErrAccountNotFound) {
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

	_, err = auc.redis.ZAdd(ctx, "transfers", redis.Z{Member: id}).Result()
	if err != nil {
		l.Error("error add id account in redis:", zap.Error(err))
	}

	return id, nil
}
