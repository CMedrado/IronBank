package transfer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

type UseCase struct {
	StoredTransfer Repository
	redis          *redis.Client
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(accountOrigin entities.Account, accountToken entities.Token, token string) ([]entities.Transfer, error) {
	err := CheckToken(token, accountToken)
	if err != nil {
		return []entities.Transfer{}, err

	}

	err = domain.CheckExistID(accountOrigin)
	if err != nil {
		return []entities.Transfer{}, err

	}

	transfers, err := auc.StoredTransfer.ReturnTransfer(accountOrigin.ID)
	if err != nil {
		return []entities.Transfer{}, domain.ErrSelect
	}

	return transfers, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(ctx context.Context, accountOriginID uuid.UUID, accountToken entities.Token, token string, accountOrigin entities.Account, accountDestination entities.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entities.Account, entities.Account) {
	err := CheckAmount(amount)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckToken(token, accountToken)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckCompareID(accountOriginID, accountDestinationIdUUID)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = domain.CheckExistID(accountDestination)
	if err != nil {
		return ErrInvalidDestinationID, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	id, _ := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := entities.Transfer{ID: id, OriginAccountID: accountOriginID, DestinationAccountID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}

	err = auc.StoredTransfer.SaveTransfer(transfer)
	if err != nil {
		return domain.ErrInsert, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	if err = auc.redis.PFAdd(ctx, "transfers_statistic", fmt.Sprint(transfer.ID)).Err(); err != nil {
		return fmt.Errorf("error no pf add: %w", err), uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	if err = auc.redis.ZIncrBy(ctx, "transfers_rank", 1, fmt.Sprint(accountOriginID)).Err(); err != nil {
		return fmt.Errorf("error no zincrby: %w", err), uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	return nil, id, accountOrigin, accountDestination
}

func (auc *UseCase) GetStatisticTransfer(ctx context.Context) (int64, error) {
	statistic, err := auc.redis.PFCount(ctx, "transfers_statistic").Result()
	if err != nil {
		return 0, domain.ErrGetRedis
	}

	return statistic, nil
}

func (auc *UseCase) GetRankTransfer(ctx context.Context) ([]string, error) {
	res14, err := auc.redis.ZRange(ctx, "transfers_rank", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return res14, nil
}

func NewUseCase(repository Repository, redis *redis.Client) *UseCase {
	return &UseCase{StoredTransfer: repository, redis: redis}
}
