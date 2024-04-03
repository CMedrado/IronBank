package authentication

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain"
)

// DecoderToken returns the ID that was inside an encrypted code
func DecoderToken(token string) (uuid.UUID, uuid.UUID, error) {
	tokeDecode, _ := base64.StdEncoding.DecodeString(token)
	idString := strings.Split(string(tokeDecode), ":")

	idInt, err := uuid.Parse(idString[3])
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, domain.ErrInvalidToken
	}

	idToken, err := uuid.Parse(idString[4])
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, domain.ErrInvalidToken
	}

	return idInt, idToken, nil
}
