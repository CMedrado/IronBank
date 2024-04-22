package transfer

import "context"

func (auc UseCase) GetCountTransfer(ctx context.Context) (int64, error) {
	statistic, err := auc.redis.PFCount(ctx, "transfers_statistic").Result()
	if err != nil {
		return 0, err
	}

	return statistic, nil
}
