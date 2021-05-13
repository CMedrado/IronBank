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

	for _, a := range transfers {
		transfer = append(transfer, a)
	}

	return transfer, nil
}

func (auc AccountUsecase) MakeTransfers(accountOriginID int, token int, accountDestinationID int, amount int) (error, int) {
	tokenStore := auc.Token.GetTokenID(accountOriginID)

	err := CheckID(token, accountOriginID, tokenStore)
	if err != nil {
		return err, 0
	}

	accountOrigin := auc.SearchID(accountOriginID)
	accountDestination := auc.SearchID(accountDestinationID)

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
	transfer := store.Transfer{id, accountOriginID, accountDestinationID, amount, CreatedAt()}
	auc.Transfer.CreatedTransfer(transfer)

	return nil, id
}
