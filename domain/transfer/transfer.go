package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
)

type UseCase struct {
	AccountUseCase domain.AccountUseCase
	TokenUseCase   domain.LoginUseCase
	StoredTransfer Repository
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(token string) ([]domain.Transfer, error) {
	var transfer []domain.Transfer
	accountOriginID := DecoderToken(token)
	transfers := auc.StoredTransfer.ReturnTransfers()
	accountToken := auc.TokenUseCase.GetTokenID(accountOriginID)

	err := domain.CheckToken(token, accountToken)

	if err != nil {
		return transfer, err
	}

	for _, a := range transfers {
		if a.AccountOriginID == accountOriginID {
			transfer = append(transfer, ChangeTransferStorage(a))
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
	transfer := domain.Transfer{ID: id, AccountOriginID: accountOriginID, AccountDestinationID: accountDestinationID, Amount: amount, CreatedAt: createdAt}
	auc.StoredTransfer.SaveTransfers(ChangeTransferDomain(transfer))

	return nil, id
}
