package transfer

var accountTransferID = make(map[int]Transfer)
var accountTransferAccountID = make(map[int]StoredTransferID)

type Transfer struct {
	ID                   int    `json:"id"`
	AccountOriginID      int    `json:"account_origin_id"`
	AccountDestinationID int    `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}

type StoredTransferID struct {
	accountTransferID map[int]Transfer
}

type StoredTransferAccountID struct {
	accountTransferAccountID map[int]StoredTransferID
}

func NewStoredTransferID() *StoredTransferID {
	return &StoredTransferID{accountTransferID}
}

func NewStoredTransferAccountID() *StoredTransferAccountID {
	return &StoredTransferAccountID{accountTransferAccountID}
}
