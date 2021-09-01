package token

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Storage) ReturnTokenID(id uuid.UUID) (entities.Token, error) {
	var token entities.Token
	statement := `SELECT * FROM tokens WHERE id_token=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&token.ID, &token.IdAccount, &token.CreatedAt)
	if err != nil {
		a.log.WithFields(log.Fields{
			"module": "returnTokenID",
			"type":   http.StatusInternalServerError,
			"time":   domain.CreatedAt(),
		}).Error(err)
		return entities.Token{}, err
	}
	return token, nil
}
