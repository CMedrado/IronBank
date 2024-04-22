package account

import (
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	StoredAccount Repository
	redis         *redis.Client
}

func NewUseCase(repository Repository, redis *redis.Client) UseCase {
	return UseCase{StoredAccount: repository, redis: redis}
}
