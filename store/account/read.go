package account

import (
	"encoding/json"
	"io"
)

type Accounts []Account

func NewAccount(rdr io.Reader) ([]Account, error) {
	var account []Account
	err := json.NewDecoder(rdr).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, err
}
