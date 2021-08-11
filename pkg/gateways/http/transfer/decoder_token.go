package transfer

import (
	"encoding/base64"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/google/uuid"
	"strings"
)

// DecoderToken returns the ID that was inside an encrypted code
func DecoderToken(token string) (uuid.UUID, uuid.UUID, error) {
	tokeDecode, _ := base64.StdEncoding.DecodeString(token)
	idString := strings.Split(string(tokeDecode), ":")
	idInt, err := uuid.Parse(idString[3])
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, domain2.ErrInvalidToken
	}
	idToken, err := uuid.Parse(idString[4])
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, domain2.ErrInvalidToken
	}
	return idInt, idToken, nil
}
