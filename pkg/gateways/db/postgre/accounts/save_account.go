package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Storage) SaveAccount(account entities.Account) error {
	statement := `INSERT INTO accounts(id, name, cpf, secret, balance, created_at)
				  VALUES ($1, $2, $3, $4, $5, $6)`
	comand, err := a.pool.Exec(context.Background(), statement, account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)
	if comand.RowsAffected() > 0 {
		return nil
	}
	if err != nil {
		a.log.WithFields(log.Fields{
			"module": "saveAccount",
			"type":   http.StatusInternalServerError,
			"time":   domain.CreatedAt(),
		}).Error(err)
		return err
	}
	return nil
}
