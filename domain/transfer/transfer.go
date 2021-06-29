package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
)

type UseCase struct {
	AccountUseCase domain.AccountUseCase
	TokenUseCase   domain.LoginUseCase
	StoredTransfer Repository
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(token string) ([]store_transfer.Transfer, error) {
	var transfer []store_transfer.Transfer
	accountOriginID := DecoderToken(token)
	transfers := auc.StoredTransfer.GetTransfers(accountOriginID)
	accountToken := auc.TokenUseCase.GetTokenID(accountOriginID)

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
	accountOrigin := auc.AccountUseCase.SearchAccount(accountOriginID)
	accountToken := auc.TokenUseCase.GetTokenID(accountOriginID)
	err = domain.CheckToken(token, accountToken)

	if err != nil {
		return err, 0
	}

	err = domain.CheckCompareID(accountOriginID, accountDestinationID)

	if err != nil {
		return err, 0
	}

	accountDestination := auc.AccountUseCase.SearchAccount(accountDestinationID)

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

	auc.AccountUseCase.UpdateBalance(accountOrigin, accountDestination)

	id := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := store_transfer.Transfer{ID: id, AccountOriginID: accountOriginID, AccountDestinationID: accountDestinationID, Amount: amount, CreatedAt: createdAt}
	auc.StoredTransfer.PostTransferID(transfer, accountOriginID)

	return nil, id
}
