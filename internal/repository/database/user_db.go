package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserDataDb struct {
	db *sqlx.DB
}

func NewUserDataDb(db *sqlx.DB) *UserDataDb {
	return &UserDataDb{db: db}
}

func (r *UserDataDb) RegisterUserIfNotExists(ctx context.Context, telegramId int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var userId int64
	err := r.db.GetContext(ctx, &userId, `INSERT INTO users (chat_id) VALUES ($1) ON CONFLICT (chat_id) DO NOTHING RETURNING id`, telegramId)
	if err != nil {
		return 0, err
	}

	return userId, err
}

func (r *UserDataDb) GetUserIdByChatId(ctx context.Context, telegramId int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var userId int64
	err := r.db.GetContext(ctx, &userId, `SELECT id FROM users WHERE chat_id = $1`, telegramId)
	if err != nil {
		return 0, err
	}

	return userId, err
}
