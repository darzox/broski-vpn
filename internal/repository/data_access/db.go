package data_access

import (
	"github.com/darzox/broski-vpn/internal/repository"
	"github.com/darzox/broski-vpn/internal/repository/database"
	"github.com/jmoiron/sqlx"
)

type DbContext struct {
	repository.UserDataStorage
	repository.KeyDataStorage
}

func NewDb(db *sqlx.DB) *DbContext {
	return &DbContext{
		database.NewUserDataDb(db),
		database.NewKeyDataDb(db),
	}
}
