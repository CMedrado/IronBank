package login

import "github.com/CMedrado/DesafioStone/domain"

// CheckLogin Checks if the cpf and secret ar correct and returns nil if not, it returns an error
func CheckLogin(accountOrigin domain.Account, newLogin domain.Login) error {
	if accountOrigin.CPF != newLogin.CPF {
		return domain.ErrInvalidCPF
	}
	if accountOrigin.Secret != newLogin.Secret {
		return domain.ErrInvalidSecret
	}
	return nil
}
