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
	CreateUserKey(ctx context.Context, userId int64, keyId int64, accessKey string, expirationDate time.Time) (int64, error)
	GetAccessKeys(ctx context.Context, userId int64) ([]dto.AccessKey, error)
	GetExpiredKeysOutlineIds(ctx context.Context) ([]int64, error)
	GetExpiredKeysWithChatIds(ctx context.Context) ([]dto.ExpiredKeyWithChatId, error)
}

type TransactionDataStorage interface {
	CreatePaymentTransaction(ctx context.Context, userId int64, keyId int64, currency string, totalAmount int, invoicePayload string, telegramPaymentChargeID string, providerPaymentChargeID string) error
}
