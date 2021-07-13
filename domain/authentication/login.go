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
	account, err := auc.AccountUseCase.GetAccountCPF(cpf)

	if err != nil {
		return domain.ErrInsert, ""
	}

	err = CheckLogin(account, newLogin)
	if err != nil {
		return domain.ErrLogin, ""
	}

	id, err := auc.AccountUseCase.GetAccountCPF(cpf)

	if err != nil {
		return domain.ErrInsert, ""
	}

	now := domain.CreatedAt()
	aux2 := 0
	var idToken uuid.UUID
	for aux2 == 0 {
		idToken, _ = domain.Random()
		searchAccount, err := auc.AccountUseCase.SearchAccount(idToken)
		if err != nil {
			return err, ""
		}
		if (searchAccount != domain.Account{}) {
			aux2 = 0
		} else {
			aux2 = 1
		}
	}
	token := now + ":" + id.ID.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	err = auc.StoredToken.SaveToken(idToken, id.ID, now)
	if err != nil {
		return domain.ErrInsert, ""
	}
	return nil, encoded
}

func (uc UseCase) GetTokenID(id uuid.UUID) (domain.Token, error) {
	tokens, err := uc.StoredToken.ReturnTokens()
	if err != nil {
		return domain.Token{}, domain.ErrInsert
	}
	token := domain.Token{}

	for _, a := range tokens {
		if a.IdAccount == id {
			token = ChangeTokenStorage(a)
		}
	}

	return token, nil
}
