package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

func (auc AccountUsecase) GetTransfers(accountOriginID int, token int) ([]store.Transfer, error) {
	transfers := auc.Transfer.GetTransfers(accountOriginID)
	tokenStore := auc.Token.GetTokenID(accountOriginID)
	var transfer []store.Transfer

	err := CheckID(token, accountOriginID, tokenStore)
	if err != nil {
		return transfer, err
	}

	for m, a := range transfers {
		if transfers[m].AccountDestinationID == accountOriginID {
			transfer = append(transfer, a)
		}
	}

	return transfer, nil
}

func (auc AccountUsecase) MakeTransfers(accountOriginID int, token int, accountDestinationID int, amount uint) (error, int) {
	err := CheckAmount(amount)

	if err != nil {
		return err, 0
	}

	tokenStore := auc.Token.GetTokenID(accountOriginID)

	err = CheckID(token, accountOriginID, tokenStore)
	if err != nil {
		return err, 0
	}

	accountOrigin, err := auc.SearchID(accountOriginID)

	if err != nil {
		return err, 0
	}

	accountDestination, err := auc.SearchID(accountDestinationID)

	if err != nil {
		return err, 0
	}

	person1 := auc.Store.TransferredBalance(accountOrigin.CPF)
	person2 := auc.Store.TransferredBalance(accountDestination.CPF)

	err = CheckBalance(person1, amount)
	if err != nil {
		return err, 0
	}

	person1.Balance = person1.Balance - amount
	person2.Balance = person2.Balance + amount

	auc.Store.UpdateBalance(person1, person2)

	id := Random()
	createdAt := CreatedAt()
	transfer := store.Transfer{id, accountOriginID, accountDestinationID, amount, createdAt}
	auc.Transfer.CreatedTransferTwo(transfer, accountDestinationID)

	return nil, id
}
