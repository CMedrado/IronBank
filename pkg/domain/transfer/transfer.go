package transfer

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UseCase struct {
	StoredTransfer Repository
	logger         *logrus.Entry
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(accountOrigin entities.Account, accountToken entities.Token, token string) ([]entities.Transfer, error) {
	l := auc.logger.WithFields(logrus.Fields{
		"module": "getTransfers",
	})
	err := CheckToken(token, accountToken)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "checkToken",
		}).Error(err)
		return []entities.Transfer{}, err

	}

	err = domain2.CheckAccountExistence(accountOrigin)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":  http.StatusBadRequest,
			"time":  domain2.CreatedAt(),
			"token": token,
			"where": "checkAccountExistence",
		}).Error(err)
		return []entities.Transfer{}, err

	}

	transfers, err := auc.StoredTransfer.ReturnTransfer(accountOrigin.ID)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":            http.StatusBadRequest,
			"time":            domain2.CreatedAt(),
			"accountOriginID": accountOrigin.ID,
			"where":           "returnTransfer",
		}).Error(err)
		return []entities.Transfer{}, domain2.ErrSelect
	}
	return transfers, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(accountOriginID uuid.UUID, accountToken entities.Token, token string, accountOrigin entities.Account, accountDestination entities.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entities.Account, entities.Account) {
	l := auc.logger.WithFields(logrus.Fields{
		"module": "getTransfers",
	})
	err := CheckAmount(amount)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":   http.StatusBadRequest,
			"time":   domain2.CreatedAt(),
			"amount": amount,
			"where":  "checkAmount",
		}).Error(err)
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckToken(token, accountToken)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":         http.StatusBadRequest,
			"time":         domain2.CreatedAt(),
			"token":        token,
			"accountToken": accountToken.CreatedAt.Format("02/01/2006 15:04:05") + ":" + accountToken.IdAccount.String() + ":" + accountToken.ID.String(),
			"where":        "checkToken",
		}).Error(err)
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckCompareID(accountOriginID, accountDestinationIdUUID)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":                 http.StatusBadRequest,
			"time":                 domain2.CreatedAt(),
			"accountOriginId":      accountOriginID,
			"accountDestinationId": accountDestinationIdUUID,
			"where":                "checkCompareID",
		}).Error(err)
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":   http.StatusBadRequest,
			"time":   domain2.CreatedAt(),
			"amount": amount,
			"where":  "checkAccountBalance",
		}).Error(err)
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = domain2.CheckExistID(accountDestination)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":                 http.StatusBadRequest,
			"time":                 domain2.CreatedAt(),
			"accountDestinationId": accountDestinationIdUUID,
			"where":                "checkExistID",
		}).Error(err)
		return ErrInvalidDestinationID, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	id, _ := domain2.Random()
	createdAt := domain2.CreatedAt()
	transfer := entities.Transfer{ID: id, OriginAccountID: accountOriginID, DestinationAccountID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}
	err = auc.StoredTransfer.SaveTransfer(transfer)
	if err != nil {
		l.WithFields(logrus.Fields{
			"type":  http.StatusInternalServerError,
			"time":  domain2.CreatedAt(),
			"where": "saveTransfer",
		}).Error(err)
		return domain2.ErrInsert, uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	return nil, id, accountOrigin, accountDestination
}

func NewUseCase(repository Repository, log *logrus.Entry) *UseCase {
	return &UseCase{StoredTransfer: repository, logger: log}
}
