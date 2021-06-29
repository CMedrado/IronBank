package transfer

import (
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
)

// DecoderToken returns the ID that was inside an encrypted code
func DecoderToken(token string) uuid.UUID {
	tokeDecode, _ := base64.StdEncoding.DecodeString(token)
	idString := strings.Split(string(tokeDecode), ":")
	idInt := uuid.MustParse(idString[3])

	return idInt
}
