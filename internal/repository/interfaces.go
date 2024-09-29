package repository

import (
	"context"
	"time"
)

type UserDataStorage interface {
	RegisterUserIfNotExists(ctx context.Context, telegramId int64) (int64, error)
}

type KeyDataStorage interface {
	CreateUserKey(ctx context.Context, userId int64, keyId int64, accessKey string, expirationDate time.Time) error
}
