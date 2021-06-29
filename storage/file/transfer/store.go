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

type StoredTransferAccountID struct {
	dataBase  io.ReadWriteSeeker
	transfers Transfers
}

func NewStoredTransferAccountID(dataBase io.ReadWriteSeeker) *StoredTransferAccountID {
	dataBase.Seek(0, 0)
	transfers, _ := NewTransfer(dataBase)

	return &StoredTransferAccountID{dataBase: dataBase, transfers: transfers}
}
