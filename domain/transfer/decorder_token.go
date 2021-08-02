package transfer

import (
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
)

// DecoderToken returns the ID that was inside an encrypted code
func DecoderToken(token string) (uuid.UUID, uuid.UUID, error) {
	tokeDecode, _ := base64.StdEncoding.DecodeString(token)
	idString := strings.Split(string(tokeDecode), ":")
	idInt, err := uuid.Parse(idString[3])
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	idToken, err := uuid.Parse(idString[4])
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	return idInt, idToken, nil
}
