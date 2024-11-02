package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/darzox/broski-vpn/internal/dto"
	"github.com/jmoiron/sqlx"
)

type KeyDataDb struct {
	db *sqlx.DB
}

func NewKeyDataDb(db *sqlx.DB) *KeyDataDb {
	return &KeyDataDb{db: db}
}

func (k *KeyDataDb) CreateUserKey(ctx context.Context, userId int64, keyId int64, accessKey string, expirationDate time.Time) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var userKeyId int64

	err := k.db.GetContext(ctx, &userKeyId, `INSERT INTO users_keys (user_id, key_id, access_key, expiration_date) VALUES ($1, $2, $3, $4) RETURNING id`, userId, keyId, accessKey, expirationDate)

	return userKeyId, err
}

func (k *KeyDataDb) GetAccessKeys(ctx context.Context, userId int64) ([]dto.AccessKey, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var accessKeys []dto.AccessKey
	err := k.db.SelectContext(ctx, &accessKeys, `SELECT access_key, expiration_date FROM users_keys WHERE user_id = $1 and expiration_date > now()`, userId)
	if err != nil {
		return nil, err
	}

	if len(accessKeys) == 0 {
		return nil, sql.ErrNoRows
	}

	return accessKeys, err
}

func (k *KeyDataDb) GetExpiredKeysOutlineIds(ctx context.Context) ([]int64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var ids []int64
	err := k.db.SelectContext(ctx, &ids, `SELECT key_id FROM users_keys WHERE expiration_date < now()`)
	if err != nil {
		return nil, err
	}

	return ids, err
}

func (k *KeyDataDb) GetExpiredKeysWithChatIds(ctx context.Context) ([]dto.ExpiredKeyWithChatId, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var expiredKeys []dto.ExpiredKeyWithChatId
	err := k.db.SelectContext(ctx, &expiredKeys, `SELECT 
	uk.key_id as key_id, u.chat_id as chat_id 
	FROM users_keys uk left join users u ON uk.user_id = u.id WHERE uk.expiration_date < now()`)
	if err != nil {
		return nil, err
	}

	return expiredKeys, err
}
