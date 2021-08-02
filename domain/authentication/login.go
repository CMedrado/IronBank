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
	id, err := auc.AccountUseCase.GetAccountCPF(cpf)

	if err != nil {
		return domain.ErrInsert, ""
	}

	err = CheckLogin(id, newLogin)
	if err != nil {
		return domain.ErrLogin, ""
	}

	now := domain.CreatedAt()
	idToken, _ := domain.Random()
	token := now.Format("02/01/2006 15:04:05") + ":" + id.ID.String() + ":" + idToken.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	save := domain.Token{ID: idToken, IdAccount: id.ID, CreatedAt: now}
	err = auc.StoredToken.SaveToken(save)
	if err != nil {
		return domain.ErrInsert, ""
	}
	return nil, encoded
}

func (uc UseCase) GetTokenID(id uuid.UUID) (domain.Token, error) {
	token, err := uc.StoredToken.ReturnTokenID(id)
	if err != nil {
		return domain.Token{}, domain.ErrInsert
	}
	return token, nil
}
