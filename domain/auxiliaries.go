package domain

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/CMedrado/DesafioStone/store"
	"math/rand"
	"strings"
	"time"
)

func CreatedAt() string {
	return time.Now().Format("02/01/2006 03:03:05")
}

func Random() int {
	return rand.Intn(100000000)
}

func (auc AccountUsecase) SearchID(id int) (store.Account, error) {
	accounts := auc.Store.TransferredAccounts()
	account := store.Account{}

	for _, a := range accounts {
		if a.ID == id {
			account = a
		}
	}

	err := CheckExistID(account)

	return account, err
}

func CpfReplace(cpf string) string {
	cpfReplace := strings.Replace(cpf, ".", "", 2)
	cpfReplace = strings.Replace(cpfReplace, "-", "", 1)
	return cpfReplace
}

func Hash(secret string) string {
	secretHash := md5.New()
	secretHash.Write([]byte(secret))
	secretHashFinal := hex.EncodeToString(secretHash.Sum(nil))
	return secretHashFinal
}
