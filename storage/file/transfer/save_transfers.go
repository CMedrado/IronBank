package transfer

import (
	"encoding/json"
)

func (a *StoredTransferAccount) SaveTransfers(transfer Transfer) {
	a.transfers = append(a.transfers, transfer)
	a.dataBase.Seek(0, 0)

	json.NewEncoder(a.dataBase).Encode(a.transfers)
}
