package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	StoredTransfer Repository
	logger         *logrus.Entry
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(accountOrigin entities.Account, accountToken entities.Token, token string) ([]entities.Transfer, error) {
	err := CheckToken(token, accountToken)
	if err != nil {
		return []entities.Transfer{}, err

	}

	err = domain.CheckExistID(accountOrigin)
	if err != nil {
		return []entities.Transfer{}, err

	}

	transfers, err := auc.StoredTransfer.ReturnTransfer(accountOrigin.ID)
	if err != nil {
		return []entities.Transfer{}, domain.ErrSelect
	}
	return transfers, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(accountOriginID uuid.UUID, accountToken entities.Token, token string, accountOrigin entities.Account, accountDestination entities.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entities.Account, entities.Account) {
	err := CheckAmount(amount)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckToken(token, accountToken)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckCompareID(accountOriginID, accountDestinationIdUUID)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = domain.CheckExistID(accountDestination)
	if err != nil {
		return ErrInvalidDestinationID, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	err = CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, uuid.UUID{}, entities.Account{}, entities.Account{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	id, _ := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := entities.Transfer{ID: id, OriginAccountID: accountOriginID, DestinationAccountID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}
	err = auc.StoredTransfer.SaveTransfer(transfer)
	if err != nil {
		return domain.ErrInsert, uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	return nil, id, accountOrigin, accountDestination
}

func NewUseCase(repository Repository, log *logrus.Entry) *UseCase {
	return &UseCase{StoredTransfer: repository, logger: log}
}
