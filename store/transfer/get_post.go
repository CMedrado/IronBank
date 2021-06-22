package transfer

func (a *StoredTransferAccountID) GetTransfers(accountOriginID int) map[int]Transfer {
	return accountTransferAccountID[accountOriginID].accountTransferID
}

func (a *StoredTransferID) PostTransferAccountID(transfer Transfer) StoredTransferID {
	accountTransferID[transfer.ID] = transfer
	return StoredTransferID{accountTransferID: accountTransferID}
}

func (a *StoredTransferAccountID) PostTransferID(transfer Transfer, id int) {
	storedTransfer := StoredTransferID{}
	transferAccount := storedTransfer.PostTransferAccountID(transfer)
	accountTransferAccountID[id] = transferAccount
}
