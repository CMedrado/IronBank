package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type UseCase struct {
	AccountUseCase domain.AccountUseCase
	TokenUseCase   domain.LoginUseCase
	StoredTransfer Repository
}

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(token string) ([]domain.Transfer, error) {
	var transfer []domain.Transfer
	accountOriginID := DecoderToken(token)
	transfers := auc.StoredTransfer.ReturnTransfers()
	accountToken := auc.TokenUseCase.GetTokenID(accountOriginID)

	err := domain.CheckToken(token, accountToken)

	if err != nil {
		return transfer, err
	}

	for _, a := range transfers {
		if a.AccountOriginID == accountOriginID {
			transfer = append(transfer, ChangeTransferStorage(a))
		}
	}

	return transfer, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(token string, accountDestinationID string, amount int) (error, uuid.UUID) {
	err := domain.CheckAmount(amount)
	accountDestinationIdUUID := uuid.MustParse(accountDestinationID)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountOriginID := DecoderToken(token)
	accountOrigin := auc.AccountUseCase.SearchAccount(accountOriginID)
	accountToken := auc.TokenUseCase.GetTokenID(accountOriginID)
	err = domain.CheckToken(token, accountToken)

	if err != nil {
		return err, uuid.UUID{}
	}

	err = domain.CheckCompareID(accountOriginID, accountDestinationIdUUID)

	if err != nil {
		return err, uuid.UUID{}
	}

	accountDestination := auc.AccountUseCase.SearchAccount(accountDestinationIdUUID)

	err = domain.CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = domain.CheckExistDestinationID(accountDestination)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	auc.AccountUseCase.UpdateBalance(accountOrigin, accountDestination)

	id, _ := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := domain.Transfer{ID: id, AccountOriginID: accountOriginID, AccountDestinationID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}
	auc.StoredTransfer.SaveTransfers(ChangeTransferDomain(transfer))

	return nil, id
}
