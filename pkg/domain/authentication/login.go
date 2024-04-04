package authentication

import (
	"encoding/base64"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

type UseCase struct {
	StoredToken Repository
}

// AuthenticatedLogin authenticates the account and returns a token
func (auc UseCase) AuthenticatedLogin(secret string, account entities.Account) (error, string) {
	secretHash := domain.CreateHash(secret)
	now := domain.CreatedAt()
	idToken, _ := domain.Random()
	newLogin := entities.Login{CPF: account.CPF, Secret: secretHash}

	err := CheckLogin(account, newLogin)
	if err != nil {
		return ErrLogin, ""
	}

	token := now.Format("02/01/2006 15:04:05") + ":" + account.ID.String() + ":" + idToken.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	save := entities.Token{ID: idToken, IdAccount: account.ID, CreatedAt: now}

	err = auc.StoredToken.SaveToken(save)
	if err != nil {
		return domain.ErrInsert, ""
	}

	return nil, encoded
}

func (auc UseCase) GetTokenID(id uuid.UUID) (entities.Token, error) {
	token, err := auc.StoredToken.ReturnTokenID(id)
	if err != nil {
		return entities.Token{}, domain.ErrInsert
	}
	return token, nil
}

func NewUseCase(repository Repository) *UseCase {
	return &UseCase{StoredToken: repository}
}
