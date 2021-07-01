package authentication

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type UseCase struct {
	AccountUseCase domain.AccountUseCase
	StoredToken    Repository
}

// AuthenticatedLogin authenticates the account and returns a token
func (auc UseCase) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := domain.CreateHash(secret)

	err, cpf := domain.CheckCPF(cpf)
	if err != nil {
		return domain.ErrLogin, ""
	}

	newLogin := domain.Login{CPF: cpf, Secret: secretHash}
	account := auc.AccountUseCase.GetAccountCPF(cpf)

	err = CheckLogin(account, newLogin)
	if err != nil {
		return domain.ErrLogin, ""
	}

	id := auc.AccountUseCase.GetAccountCPF(cpf)
	now := domain.CreatedAt()
	token := now + ":" + id.ID.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	auc.StoredToken.SaveToken(id.ID, encoded)

	return nil, encoded
}

func (uc UseCase) GetTokenID(id uuid.UUID) domain.Token {
	tokens := uc.StoredToken.ReturnTokens()
	token := domain.Token{}

	for _, a := range tokens {
		if a.ID == id {
			token = ChangeTokenStorage(a)
		}
	}

	return token
}
