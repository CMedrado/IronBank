package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Storage) ReturnAccountID(id uuid.UUID) (entities.Account, error) {
	var account entities.Account
	statement := `SELECT * FROM accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil && err.Error() != ("no rows in result set") {
		a.log.WithFields(log.Fields{
			"module": "returnAccountID",
			"type":   http.StatusInternalServerError,
			"time":   domain.CreatedAt(),
		}).Error(err)
		return entities.Account{}, err
	}
	return account, nil
}
