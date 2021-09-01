package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Storage) ReturnAccountCPF(cpf string) (entities.Account, error) {

	var account entities.Account
	statement := `SELECT * FROM accounts WHERE cpf=$1`
	err := a.pool.QueryRow(context.Background(), statement, cpf).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil && err.Error() != ("no rows in result set") {
		a.log.WithFields(log.Fields{
			"module": "returnAccountCPF",
			"type":   http.StatusInternalServerError,
			"time":   domain.CreatedAt(),
		}).Error(err)
		return entities.Account{}, err
	}
	return account, nil
}
