package job

import (
	"context"
	"log/slog"
)

type Usecase interface {
	RemoveExpiredKeys(ctx context.Context) error
}

type cronjob struct {
	usecase Usecase
	logger  *slog.Logger
}

func New(logger *slog.Logger, usecase Usecase) *cronjob {
	return &cronjob{
		usecase: usecase,
		logger:  logger,
	}
}

func (j *cronjob) RemoveExpiredKeys() {
	j.logger.Info("started check for expired keys to remove")
	err := j.usecase.RemoveExpiredKeys(context.Background())
	if err != nil {
		j.logger.Error("failed to remove expired keys")
	}
	j.logger.Info("finished check for expired keys to remove")
}
