package transfer

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type UseCase struct {
	StoredTransfer Repository
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(accountOriginID uuid.UUID, accountToken entries.Token, token string) ([]entries.Transfer, error) {
	err := CheckToken(token, accountToken)
	if err != nil {
		return []entries.Transfer{}, err
	}
	transfers, err := auc.StoredTransfer.ReturnTransfer(accountOriginID)
	if err != nil {
		return []entries.Transfer{}, domain2.ErrSelect
	}
	return transfers, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(accountOriginID uuid.UUID, accountToken entries.Token, token string, accountOrigin entries.Account, accountDestination entries.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entries.Account, entries.Account) {
	err := CheckAmount(amount)
	if err != nil {
		return err, uuid.UUID{}, entries.Account{}, entries.Account{}
	}

	err = CheckToken(token, accountToken)
	if err != nil {
		return err, uuid.UUID{}, entries.Account{}, entries.Account{}
	}

	err = CheckCompareID(accountOriginID, accountDestinationIdUUID)
	if err != nil {
		return err, uuid.UUID{}, entries.Account{}, entries.Account{}
	}

	err = CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, uuid.UUID{}, entries.Account{}, entries.Account{}
	}

	err = domain2.CheckExistID(accountDestination)
	if err != nil {
		return domain2.ErrInvalidDestinationID, uuid.UUID{}, entries.Account{}, entries.Account{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	id, _ := domain2.Random()
	createdAt := domain2.CreatedAt()
	transfer := entries.Transfer{ID: id, OriginAccountID: accountOriginID, DestinationAccountID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}
	err = auc.StoredTransfer.SaveTransfer(transfer)
	if err != nil {
		return domain2.ErrInsert, uuid.UUID{}, entries.Account{}, entries.Account{}
	}
	return nil, id, accountOrigin, accountDestination
}
