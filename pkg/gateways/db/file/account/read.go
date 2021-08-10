package account

import (
	"encoding/json"
	"io"
)

func NewAccount(rdr io.Reader) ([]Account, error) {
	var account []Account
	err := json.NewDecoder(rdr).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, err
}

func (accounts Accounts) Find(name string) *Account {
	for i, p := range accounts {
		if p.Name == name {
			return &accounts[i]
		}
	}
	return nil
}
