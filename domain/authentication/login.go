package authentication

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/domain"
	token2 "github.com/CMedrado/DesafioStone/storage/postgre/token"
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
	idToken, _ := domain.Random()
	token := now.Format("02/01/2006 15:04:05") + ":" + id.ID.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	save := token2.Token{ID: idToken, IdAccount: id.ID, CreatedAt: now}
	err = auc.StoredToken.SaveToken(save)
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

func (auc UseCase) SearchToken(id uuid.UUID) (domain.Token, error) {
	tokens, err := auc.StoredToken.ReturnTokens()
	if err != nil {
		return domain.Token{}, domain.ErrInsert
	}
	token := domain.Token{}

	for _, a := range tokens {
		if a.ID == id {
			token = ChangeTokenStorage(a)
		}
	}

	return token, nil
}
