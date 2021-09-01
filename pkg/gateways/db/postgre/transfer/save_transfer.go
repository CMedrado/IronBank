package transfer

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Storage) SaveTransfer(transfer entities.Transfer) error {
	statement := `INSERT INTO transfers(id, origin_account_id, destination_account_id, amount, created_at)
				  VALUES ($1, $2, $3, $4, $5)`
	comand, err := a.pool.Exec(context.Background(), statement, transfer.ID, transfer.OriginAccountID, transfer.DestinationAccountID, transfer.Amount, transfer.CreatedAt)
	if comand.RowsAffected() > 0 {
		return nil
	}
	if err != nil {
		a.log.WithFields(log.Fields{
			"module": "saveTransfer",
			"type":   http.StatusInternalServerError,
			"time":   domain.CreatedAt(),
		}).Error(err)
		return err
	}
	return nil
}
