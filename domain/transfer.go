package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

func GetTransfers(id int) []store.Transfer {
	transferMethod := store.StoredTransfer{}
	transfers := transferMethod.GetTransfers()
	var transfer []store.Transfer
	transerr := transfers[id]
	for _, a := range transfers[id] {
		transfer = append(transfer, a)
	}

	return transfer
}

func MakeTransfers(accountOriginID int, cpfDestination string, amount int) error {
	accountsMethods := store.StoredAccount{}
	loginMethod := store.StoredLogin{}
	loginTransfer := store.StoredTransfer{}
	login := loginMethod.GetLogin(accountOriginID)
	person1 := accountsMethods.TransferredBalance(login.CPF)
	person2 := accountsMethods.TransferredBalance(cpfDestination)
	err := CheckLogin(login)
	if err != nil {
		return err
	}
	err = CheckBalance(person1, amount)
	if err != nil {
		return err
	}
	person1.Balance = person1.Balance - amount
	person2.Balance = person2.Balance + amount
	accountsMethods.UpdateBalance(person1, person2)
	accountOriginID = Random()
	createdAt := CreatedAt()
	transfer := store.Transfer{person1.ID, accountOriginID, cpfDestination, amount, createdAt}
	loginTransfer.CreatedTransfer(transfer)
	return nil
}
