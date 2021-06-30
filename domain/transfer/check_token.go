package transfer

import "github.com/CMedrado/DesafioStone/domain"

// CheckToken checks if the token is correct and returns nil if not, it returns an error
func CheckToken(token string, tokens domain.Token) error {
	if token != tokens.Token {
		return domain.ErrInvalidToken
	}
	return nil
}
