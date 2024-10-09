package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransactionDataDb struct {
	db *sqlx.DB
}

func NewTransactionDataDb(db *sqlx.DB) *TransactionDataDb {
	return &TransactionDataDb{db: db}
}

func (r *TransactionDataDb) CreatePaymentTransaction(ctx context.Context, userId int64, keyId int64, currency string, totalAmount int, invoicePayload string, telegramPaymentChargeID string, providerPaymentChargeID string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.ExecContext(ctx, `INSERT INTO transactions (user_id, key_id, currency, amount, invoice_payload, telegram_payment_charge_id, provider_payment_charge_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`, userId, keyId, currency, totalAmount, invoicePayload, telegramPaymentChargeID, providerPaymentChargeID)

	return err
}
