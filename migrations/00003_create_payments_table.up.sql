BEGIN;

create table if not exists payments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL, -- pending, completed, failed
    transaction_id VARCHAR(100),
    payment_method VARCHAR(50), -- Например, 'credit_card', 'crypto'
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;