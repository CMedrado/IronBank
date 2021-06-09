package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/store"
)

type UseCase struct {
	StoredAccount  *store.StoredAccount
	StoredToken    *store.StoredToken
	StoredTransfer *store.StoredTransferAccountID
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(token string) ([]store.Transfer, error) {
	var transfer []store.Transfer
	accountOriginID := domain.DecoderToken(token)
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
	accountUseCase := account.UseCase{StoredAccount: auc.StoredAccount}

	if err != nil {
		return err, 0
	}

	accountOriginID := domain.DecoderToken(token)
	accountOrigin := accountUseCase.SearchAccount(accountOriginID)
	accountToken := auc.StoredToken.GetTokenID(accountOriginID)
	err = domain.CheckToken(token, accountToken)

	if err != nil {
		return err, 0
	}

	err = domain.CheckCompareID(accountOriginID, accountDestinationID)

	if err != nil {
		return err, 0
	}

	accountDestination := accountUseCase.SearchAccount(accountDestinationID)

	person1 := auc.StoredAccount.GetBalance(accountOrigin.CPF)
	person2 := auc.StoredAccount.GetBalance(accountDestination.CPF)

	err = domain.CheckAccountBalance(person1, amount)
	if err != nil {
		return err, 0
	}

	err = domain.CheckExistDestinationID(accountDestination)
	if err != nil {
		return err, 0
	}

	person1.Balance = person1.Balance - amount
	person2.Balance = person2.Balance + amount

	auc.StoredAccount.UpdateBalance(person1, person2)

	id := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := store.Transfer{ID: id, AccountOriginID: accountOriginID, AccountDestinationID: accountDestinationID, Amount: amount, CreatedAt: createdAt}
	auc.StoredTransfer.PostTransferID(transfer, accountOriginID)

	return nil, id
}
