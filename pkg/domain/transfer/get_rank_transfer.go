package transfer

import "context"

func (auc UseCase) GetRankTransfer(ctx context.Context) ([]string, error) {
	res14, err := auc.redis.ZRange(ctx, "transfers_rank", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return res14, nil
}
