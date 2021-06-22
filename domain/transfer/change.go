package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/storage/file/transfer"
)

func ChangeTransferDomain(transferDomain domain.Transfer) transfer.Transfer {
	transferStorage := transfer.Transfer{ID: transferDomain.ID, AccountOriginID: transferDomain.AccountOriginID, AccountDestinationID: transferDomain.AccountDestinationID, Amount: transferDomain.Amount, CreatedAt: transferDomain.CreatedAt}
	return transferStorage
}

func ChangeTransferStorage(transferStorage transfer.Transfer) domain.Transfer {
	transferDomain := domain.Transfer{ID: transferStorage.ID, AccountOriginID: transferStorage.AccountOriginID, AccountDestinationID: transferStorage.AccountDestinationID, Amount: transferStorage.Amount, CreatedAt: transferStorage.CreatedAt}
	return transferDomain
}
