package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func Set(ctx context.Context, input entities.Account, r *goredis.Client) error {
	inputJson, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error marshal account: %w", err)
	}

	err = r.Set(ctx, input.CPF, inputJson, 30*time.Second).Err()
	if err != nil {
		return fmt.Errorf("error set account: %w", err)
	}

	return nil
}

func Get(ctx context.Context, cpf string, r *goredis.Client) (entities.Account, error) {
	result, err := r.Get(ctx, cpf).Result()
	if err != nil {
		if errors.Is(err, goredis.Nil) {
			return entities.Account{}, domain.ErrAccountNotFound
		}
		return entities.Account{}, fmt.Errorf("error retrieving cached account: %w", err)
	}

	var output entities.Account
	err = json.Unmarshal([]byte(result), &output)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error unmarshall account: %w", err)
	}

	return output, nil
}
