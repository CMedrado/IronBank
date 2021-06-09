package domain

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CreatedAt returns the current date and time
func CreatedAt() string {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	utc := time.Now().In(loc).Format("02/01/2006 15:04:05")
	return utc
}

// Random returns a random number
func Random() int {
	return rand.Intn(100000000)
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
