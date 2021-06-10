package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
)

type UseCase struct {
	StoredAccount  domain.AccountUsecase
	StoredToken    *store.StoredToken
	StoredTransfer *store.StoredTransferAccountID
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(token string) ([]store.Transfer, error) {
	var transfer []store.Transfer
	accountOriginID := DecoderToken(token)
	transfers := auc.StoredTransfer.GetTransfers(accountOriginID)
	accountToken := auc.StoredToken.GetTokenID(accountOriginID)

	err := domain.CheckToken(token, accountToken)

	if err != nil {
		return transfer, err
	}

	for m, a := range transfers {
		if transfers[m].AccountOriginID == accountOriginID {
			transfer = append(transfer, a)
		}
	}

	return transfer, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(token string, accountDestinationID int, amount int) (error, int) {
	err := domain.CheckAmount(amount)

	if err != nil {
		return err, 0
	}

	accountOriginID := DecoderToken(token)
	accountOrigin := auc.StoredAccount.SearchAccount(accountOriginID)
	accountToken := auc.StoredToken.GetTokenID(accountOriginID)
	err = domain.CheckToken(token, accountToken)

	if err != nil {
		return err, 0
	}

	err = domain.CheckCompareID(accountOriginID, accountDestinationID)

	if err != nil {
		return err, 0
	}

	accountDestination := auc.StoredAccount.SearchAccount(accountDestinationID)

	err = domain.CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, 0
	}

	err = domain.CheckExistDestinationID(accountDestination)
	if err != nil {
		return err, 0
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	auc.StoredAccount.UpdateBalance(accountOrigin, accountDestination)

	id := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := store.Transfer{ID: id, AccountOriginID: accountOriginID, AccountDestinationID: accountDestinationID, Amount: amount, CreatedAt: createdAt}
	auc.StoredTransfer.PostTransferID(transfer, accountOriginID)

	return nil, id
}
