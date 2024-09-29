package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type KeyDataDb struct {
	db *sqlx.DB
}

func NewKeyDataDb(db *sqlx.DB) *KeyDataDb {
	return &KeyDataDb{db: db}
}

func (k *KeyDataDb) CreateUserKey(ctx context.Context, userId int64, keyId int64, accessKey string, expirationDate time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := k.db.ExecContext(ctx, `INSERT INTO users_keys (user_id, key_id, access_key, expiration_date) VALUES ($1, $2, $3, $4)`, userId, keyId, accessKey, expirationDate)

	return err
}
