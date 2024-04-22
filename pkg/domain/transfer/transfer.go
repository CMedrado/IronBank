package transfer

import (
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	StoredTransfer Repository
	redis          *redis.Client
}

func NewUseCase(repository Repository, redis *redis.Client) UseCase {
	return UseCase{StoredTransfer: repository, redis: redis}
}
