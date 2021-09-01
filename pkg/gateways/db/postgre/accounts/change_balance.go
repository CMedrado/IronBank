package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Storage) ChangeBalance(personDomain1, personDomain2 entities.Account) error {
	l := a.log.WithFields(log.Fields{
		"module": "changeBalance",
	})
	person1 := ChangeAccountDomain(personDomain1)
	person2 := ChangeAccountDomain(personDomain2)
	statement := `UPDATE accounts
				  SET balance=$1
				  WHERE name=$2`
	_, err := a.pool.Exec(context.Background(), statement, person1.Balance, person1.Name)
	if err != nil {
		l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
			"time": domain.CreatedAt(),
		}).Error(err)
		return err
	}
	a.pool.Exec(context.Background(), statement, person2.Balance, person2.Name)
	if err != nil {
		l.WithFields(log.Fields{
			"type": http.StatusInternalServerError,
			"time": domain.CreatedAt(),
		}).Error(err)
		return err
	}
	return nil
}
