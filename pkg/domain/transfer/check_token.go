package transfer

import (
	"encoding/base64"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

// CheckToken checks if the token is correct and returns nil if not, it returns an error
func CheckToken(token string, tokens entities.Token) error {
	tokenEncode := tokens.CreatedAt.Format("02/01/2006 15:04:05") + ":" + tokens.IdAccount.String() + ":" + tokens.ID.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(tokenEncode))
	if token != encoded {
		return domain2.ErrInvalidToken
	}
	return nil
}
