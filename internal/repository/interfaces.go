package repository

import (
	"context"
	"time"

	"github.com/darzox/broski-vpn/internal/dto"
)

type UserDataStorage interface {
	RegisterUserIfNotExists(ctx context.Context, telegramId int64) (int64, error)
	GetUserIdByChatId(ctx context.Context, telegramId int64) (int64, error)
}

type KeyDataStorage interface {
	CreateUserKey(ctx context.Context, userId int64, keyId int64, accessKey string, expirationDate time.Time) error
	GetAccessKeys(ctx context.Context, userId int64) ([]dto.AccessKey, error)
	GetExpiredKeysOutlineIds(ctx context.Context) ([]int64, error)
}
