BEGIN;

drop table if exists payments cascade;

create table if not exists transactions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    key_id BIGINT NOT NULL REFERENCES users_keys(id) ON DELETE CASCADE,
    currency TEXT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    invoice_payload TEXT NOT NULL,
    telegram_payment_charge_id TEXT NOT NULL,
    provider_payment_charge_id TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMIT;