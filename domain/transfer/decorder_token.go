package transfer

import (
	"encoding/base64"
	"strconv"
	"strings"
)

// DecoderToken returns the ID that was inside an encrypted code
func DecoderToken(token string) int {
	tokeDecode, _ := base64.StdEncoding.DecodeString(token)
	idString := strings.Split(string(tokeDecode), ":")
	idInt, _ := strconv.Atoi(idString[3])

	return idInt
}
