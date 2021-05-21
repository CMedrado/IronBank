package domain

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/CMedrado/DesafioStone/store"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CreatedAt returns the current date and time
func CreatedAt() string {
	return time.Now().Format("02/01/2006 03:03:05")
}

// Random returns a random number
func Random() int {
	return rand.Intn(100000000)
}

// SearchAccount returns the account via the received ID
func (auc AccountUseCase) SearchAccount(id int) store.Account {
	accounts := auc.Store.GetAccounts()
	account := store.Account{}

	for _, a := range accounts {
		if a.ID == id {
			account = a
		}
	}

	return account
}

// CpfReplace returns the CPF received in a single format
func CpfReplace(cpf string) string {
	cpfReplace := strings.Replace(cpf, ".", "", 2)
	cpfReplace = strings.Replace(cpfReplace, "-", "", 1)
	return cpfReplace
}

// CreateHash returns the secret received as a hash
func CreateHash(secret string) string {
	secretHash := md5.New()
	secretHash.Write([]byte(secret))
	secretHashFinal := hex.EncodeToString(secretHash.Sum(nil))
	return secretHashFinal
}

// DecoderToken returns the ID that was inside an encrypted code
func DecoderToken(token string) int {
	tokeDecode, _ := base64.StdEncoding.DecodeString(token)
	idString := strings.Split(string(tokeDecode), ":")
	idInt, _ := strconv.Atoi(idString[3])

	return idInt
}
