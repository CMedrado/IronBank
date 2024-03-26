package transfer

import (
	"github.com/google/uuid"
)

type Transfers []Transfer

type Transfer struct {
	ID                   uuid.UUID `json:"id"`
	AccountOriginID      uuid.UUID `json:"account_origin_id"`
	AccountDestinationID uuid.UUID `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            string    `json:"created_at"`
}

//type StoredTransferAccount struct {
//	dataBase  io.ReadWriteSeeker
//	transfers Transfers
//}

//func NewStoredTransfer(dataBase io.ReadWriteSeeker) *StoredTransferAccount {
//	dataBase.Seek(0, 0)
//	transfers, _ := NewTransfer(dataBase)
//
//	return &StoredTransferAccount{dataBase: dataBase, transfers: transfers}
//}
