package dto

import "time"

type AccessKey struct {
	AccessKeyString string    `db:"access_key"`
	ExpirationDate  time.Time `db:"expiration_date"`
}

type ExpiredKeyWithChatId struct {
	KeyId  int64 `db:"key_id"`
	ChatId int64 `db:"chat_id"`
}
