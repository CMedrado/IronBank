package transfer

type TransfersRequest struct {
	AccountOriginID      string `json:"account_origin_id"`
	Token                int    `json:"token"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}
