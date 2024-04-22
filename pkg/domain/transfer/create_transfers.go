package transfer

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(ctx context.Context, token string, amount int, accountDestinationId string) (error, uuid.UUID) {
	accountOriginID, tokenOriginID, err := domain.DecoderToken(token)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountOrigin, err := auc.StoredTransfer.ReturnAccountID(accountOriginID)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountToken, err := auc.StoredTransfer.ReturnTokenID(tokenOriginID)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountDestinationIdUUID, err := uuid.Parse(accountDestinationId)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountDestination, err := auc.StoredTransfer.ReturnAccountID(accountDestinationIdUUID)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = CheckAmount(amount)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = CheckToken(token, accountToken)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = CheckCompareID(accountOriginID, accountDestinationIdUUID)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = domain.CheckExistID(accountDestination)
	if err != nil {
		return ErrInvalidDestinationID, uuid.UUID{}
	}

	err = CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	id, _ := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := entities.Transfer{ID: id, OriginAccountID: accountOriginID, DestinationAccountID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}

	err = auc.StoredTransfer.SaveTransfer(transfer)
	if err != nil {
		return domain.ErrInsert, uuid.UUID{}
	}

	if err = auc.redis.PFAdd(ctx, "transfers_statistic", fmt.Sprint(transfer.ID)).Err(); err != nil {
		return fmt.Errorf("error no pf add: %w", err), uuid.UUID{}
	}

	if err = auc.redis.ZIncrBy(ctx, "transfers_rank", 1, fmt.Sprint(accountOriginID)).Err(); err != nil {
		return fmt.Errorf("error no zincrby: %w", err), uuid.UUID{}
	}

	if err = auc.StoredTransfer.ChangeBalance(accountOrigin, accountDestination); err != nil {
		return err, uuid.UUID{}
	}

	return nil, id
}
