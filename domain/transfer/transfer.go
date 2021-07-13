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
	accountOriginID, err := DecoderToken(token)
	if err != nil {
		return transfer, domain.ErrParse
	}
	transfers, err := auc.StoredTransfer.ReturnTransfer()
	if err != nil {
		return []domain.Transfer{}, err
	}
	accountToken, err := auc.TokenUseCase.GetTokenID(accountOriginID)
	if err != nil {
		return []domain.Transfer{}, err
	}

	err = CheckToken(token, accountToken)

	if err != nil {
		return transfer, err
	}

	for _, a := range transfers {
		if a.OriginAccountID == accountOriginID {
			transfer = append(transfer, ChangeTransferStorage(a))
		}
	}

	return transfer, nil
}

// CreateTransfers create and transfers, returns the id of the created transfer
func (auc UseCase) CreateTransfers(token string, accountDestinationID string, amount int) (error, uuid.UUID) {
	err := domain.CheckAmount(amount)

	if err != nil {
		return err, uuid.UUID{}
	}

	accountDestinationIdUUID, err := uuid.Parse(accountDestinationID)

	if err != nil {
		return domain.ErrParse, uuid.UUID{}
	}

	accountOriginID, err := DecoderToken(token)
	if err != nil {
		return domain.ErrParse, uuid.UUID{}
	}

	accountOrigin, err := auc.AccountUseCase.SearchAccount(accountOriginID)
	if err != nil {
		return err, uuid.UUID{}
	}

	accountToken, err := auc.TokenUseCase.GetTokenID(accountOriginID)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = CheckToken(token, accountToken)

	if err != nil {
		return err, uuid.UUID{}
	}

	err = CheckCompareID(accountOriginID, accountDestinationIdUUID)

	if err != nil {
		return err, uuid.UUID{}
	}

	accountDestination, err := auc.AccountUseCase.SearchAccount(accountDestinationIdUUID)
	if err != nil {
		return err, uuid.UUID{}
	}
	err = CheckAccountBalance(accountOrigin.Balance, amount)
	if err != nil {
		return err, uuid.UUID{}
	}

	err = domain.CheckExistID(accountDestination)
	if err != nil {
		return domain.ErrInvalidDestinationID, uuid.UUID{}
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	auc.AccountUseCase.UpdateBalance(accountOrigin, accountDestination)

	id, _ := domain.Random()
	createdAt := domain.CreatedAt()
	transfer := domain.Transfer{ID: id, OriginAccountID: accountOriginID, DestinationAccountID: accountDestinationIdUUID, Amount: amount, CreatedAt: createdAt}
	auc.StoredTransfer.SaveTransfer(ChangeTransferDomain(transfer))

	return nil, id
}
