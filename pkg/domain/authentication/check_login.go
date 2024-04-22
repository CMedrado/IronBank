package authentication

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

// CheckLogin Checks if the cpf and secret ar correct and returns nil if not, it returns an error
func CheckLogin(accountOrigin entities.Account, newLogin entities.Login) error {
	if accountOrigin.CPF != newLogin.CPF {
		return domain.ErrInvalidCPF
	}
	if accountOrigin.Secret != newLogin.Secret {
		return ErrInvalidSecret
	}
	return nil
}
