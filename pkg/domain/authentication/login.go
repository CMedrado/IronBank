package authentication

import (
	"encoding/base64"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type UseCase struct {
	StoredToken Repository
}

// AuthenticatedLogin authenticates the account and returns a token
func (auc UseCase) AuthenticatedLogin(cpf, secret string, account entries.Account) (error, string) {
	secretHash := domain2.CreateHash(secret)

	newLogin := entries.Login{CPF: cpf, Secret: secretHash}

	err := CheckLogin(account, newLogin)
	if err != nil {
		return domain2.ErrLogin, ""
	}

	now := domain2.CreatedAt()
	idToken, _ := domain2.Random()
	token := now.Format("02/01/2006 15:04:05") + ":" + account.ID.String() + ":" + idToken.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	save := entries.Token{ID: idToken, IdAccount: account.ID, CreatedAt: now}
	err = auc.StoredToken.SaveToken(save)
	if err != nil {
		return domain2.ErrInsert, ""
	}
	return nil, encoded
}

func (uc UseCase) GetTokenID(id uuid.UUID) (entries.Token, error) {
	token, err := uc.StoredToken.ReturnTokenID(id)
	if err != nil {
		return entries.Token{}, domain2.ErrInsert
	}
	return token, nil
}
