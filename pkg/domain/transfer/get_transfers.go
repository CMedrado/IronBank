package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

// GetTransfers returns all account transfers
func (auc UseCase) GetTransfers(token string) ([]entities.Transfer, error) {
	accountOriginID, tokenID, err := domain.DecoderToken(token)
	if err != nil {
		return []entities.Transfer{}, err
	}

	accountToken, err := auc.StoredTransfer.ReturnTokenID(tokenID)
	if err != nil {
		return []entities.Transfer{}, err
	}

	accountOrigin, err := auc.StoredTransfer.ReturnAccountID(accountOriginID)
	if err != nil {
		return []entities.Transfer{}, domain.ErrSelect
	}

	err = CheckToken(token, accountToken)
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
