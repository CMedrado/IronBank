package authentication

import (
	"encoding/base64"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UseCase struct {
	StoredToken Repository
	logger      *logrus.Entry
}

// AuthenticatedLogin authenticates the account and returns a token
func (auc UseCase) AuthenticatedLogin(secret string, account entities.Account) (error, string) {
	l := auc.logger.WithFields(logrus.Fields{
		"module": "authenticatedLogin",
	})

	secretHash := domain2.CreateHash(secret)

	newLogin := entities.Login{CPF: account.CPF, Secret: secretHash}

	err := CheckLogin(account, newLogin)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"where": "checkLogin",
		}).Error(err)
		return ErrLogin, ""
	}

	now := domain2.CreatedAt()
	idToken, _ := domain2.Random()
	token := now.Format("02/01/2006 15:04:05") + ":" + account.ID.String() + ":" + idToken.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	save := entities.Token{ID: idToken, IdAccount: account.ID, CreatedAt: now}
	err = auc.StoredToken.SaveToken(save)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":  http.StatusInternalServerError,
			"time":  domain2.CreatedAt(),
			"where": "saveToken",
		}).Error(err)
		return domain2.ErrInsert, ""
	}
	return nil, encoded
}

func (uc UseCase) GetTokenID(id uuid.UUID) (entities.Token, error) {
	token, err := uc.StoredToken.ReturnTokenID(id)
	if err != nil {
		return entities.Token{}, domain2.ErrInsert
	}
	return token, nil
}

func NewUseCase(repository Repository, log *logrus.Entry) *UseCase {
	return &UseCase{StoredToken: repository, logger: log}
}
