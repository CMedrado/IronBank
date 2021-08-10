package authentication

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
)

// CheckLogin Checks if the cpf and secret ar correct and returns nil if not, it returns an error
func CheckLogin(accountOrigin entries.Account, newLogin entries.Login) error {
	if accountOrigin.CPF != newLogin.CPF {
		return domain2.ErrInvalidCPF
	}
	if accountOrigin.Secret != newLogin.Secret {
		return domain2.ErrInvalidSecret
	}
	return nil
}
