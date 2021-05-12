package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"math/rand"
	"time"
)

func GetTransfers(id int) []store.Transfer {
	transferMethod := store.StoredTransfer{}
	transfers := transferMethod.GetTransfers()
	var transfer []store.Transfer

	for _, a := range transfers[id] {
		transfer = append(transfer, a)
	}

	return transfer
}

func MakeTransfers(accountOriginID int, cpfDestination string, amount int) {
	accountsMethods := store.StoredAccount{}
	loginMethod := store.StoredLogin{}
	loginTransfer := store.StoredTransfer{}
	login := loginMethod.GetLogin(accountOriginID)
	person1 := accountsMethods.TransferredBalance(login.CPF)
	person2 := accountsMethods.TransferredBalance(cpfDestination)
	person1.Balance = person1.Balance - amount
	person2.Balance = person2.Balance + amount
	accountsMethods.UpdateBalance(person1, person2)
	accountOriginID = rand.Intn(10000000)
	createdAt := time.Now().Format("02/01/2006 03:03:05")
	transfer := store.Transfer{person1.ID, accountOriginID, cpfDestination, amount, createdAt}
	loginTransfer.CreatedTransfer(transfer)
}
