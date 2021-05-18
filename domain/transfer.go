package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

func (auc AccountUsecase) GetTransfers(token string) ([]store.Transfer, error) {
	var transfer []store.Transfer
	accountOriginID := DecoderToken(token)
	transfers := auc.Transfer.GetTransfers(accountOriginID)
	_, err := auc.SearchID(accountOriginID)

	if err != nil {
		return transfer, err
	}

	accountToken := auc.Token.GetTokenID(accountOriginID)
	err = CheckToken(token, accountToken)

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

func (auc AccountUsecase) MakeTransfers(token string, accountDestinationID int, amount uint) (error, int) {
	err := CheckAmount(amount)

	if err != nil {
		return err, 0
	}

	accountOriginID := DecoderToken(token)
	accountOrigin, err := auc.SearchID(accountOriginID)

	if err != nil {
		return err, 0
	}

	accountToken := auc.Token.GetTokenID(accountOriginID)
	err = CheckToken(token, accountToken)

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
