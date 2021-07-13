package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/storage/postgre/transfer"
)

func ChangeTransferDomain(transferDomain domain.Transfer) transfer.Transfer {
	transferStorage := transfer.Transfer{ID: transferDomain.ID, OriginAccountID: transferDomain.OriginAccountID, DestinationAccountID: transferDomain.DestinationAccountID, Amount: transferDomain.Amount, CreatedAt: transferDomain.CreatedAt}
	return transferStorage
}

func ChangeTransferStorage(transferStorage transfer.Transfer) domain.Transfer {
	transferDomain := domain.Transfer{ID: transferStorage.ID, OriginAccountID: transferStorage.OriginAccountID, DestinationAccountID: transferStorage.DestinationAccountID, Amount: transferStorage.Amount, CreatedAt: transferStorage.CreatedAt}
	return transferDomain
}
