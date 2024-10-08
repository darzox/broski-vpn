package dto

import "time"

type AccessKey struct {
	AccessKeyString string    `db:"access_key"`
	ExpirationDate  time.Time `db:"expiration_date"`
}
