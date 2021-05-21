package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

// GetTransfers returns all account transfers
func (auc AccountUseCase) GetTransfers(token string) ([]store.Transfer, error) {
	var transfer []store.Transfer
	accountOriginID := DecoderToken(token)
	transfers := auc.Transfer.GetTransfers(accountOriginID)
	accountToken := auc.Token.GetTokenID(accountOriginID)

	err := CheckToken(token, accountToken)

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
func (auc AccountUseCase) CreateTransfers(token string, accountDestinationID int, amount int) (error, int) {
	err := CheckAmount(amount)

	if err != nil {
		return err, 0
	}

	accountOriginID := DecoderToken(token)
	accountOrigin := auc.SearchAccount(accountOriginID)
	accountToken := auc.Token.GetTokenID(accountOriginID)
	err = CheckToken(token, accountToken)

	if err != nil {
		return err, 0
	}

	err = CheckCompareID(accountOriginID, accountDestinationID)

	if err != nil {
		return err, 0
	}

	accountDestination := auc.SearchAccount(accountDestinationID)

	person1 := auc.Store.GetBalance(accountOrigin.CPF)
	person2 := auc.Store.GetBalance(accountDestination.CPF)

	err = CheckAccountBalance(person1, amount)
	if err != nil {
		return err, 0
	}

	CheckExistDestinationID(accountDestination)
	if err != nil {
		return err, 0
	}

	person1.Balance = person1.Balance - amount
	person2.Balance = person2.Balance + amount

	auc.Store.UpdateBalance(person1, person2)

	id := Random()
	createdAt := CreatedAt()
	transfer := store.Transfer{ID: id, AccountOriginID: accountOriginID, AccountDestinationID: accountDestinationID, Amount: amount, CreatedAt: createdAt}
	auc.Transfer.PostTransferID(transfer, accountOriginID)

	return nil, id
}
