package transfer

func (a *StoredTransferAccount) ReturnTransfers() []Transfer {
	return a.transfers
}
