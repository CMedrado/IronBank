package transfer

import (
	"io"
)

type Transfers []Transfer

type Transfer struct {
	ID                   int    `json:"id"`
	AccountOriginID      int    `json:"account_origin_id"`
	AccountDestinationID int    `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}

type StoredTransferAccount struct {
	dataBase  io.ReadWriteSeeker
	transfers Transfers
}

func NewStoredTransfer(dataBase io.ReadWriteSeeker) *StoredTransferAccount {
	dataBase.Seek(0, 0)
	transfers, _ := NewTransfer(dataBase)

	return &StoredTransferAccount{dataBase: dataBase, transfers: transfers}
}
