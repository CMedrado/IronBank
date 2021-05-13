package domain

import (
	"crypto/sha1"
	"fmt"
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

func (auc AccountUsecase) SearchID(id int) store.Account {
	accounts := auc.Store.TransferredAccounts()
	account := store.Account{}
	for _, a := range accounts {
		if a.ID == id {
			account = a
		}
	}

	return account
}

func CpfReplace(cpf string) string {
	cpfReplace := strings.Replace(cpf, ".", "", 2)
	cpfReplace = strings.Replace(cpf, "-", "", 1)
	return cpfReplace
}

func hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	sum := h.Sum(nil)
	armored := fmt.Sprintf("%x", sum)
	return armored
}
