BEGIN;

create table if not exists data_usage(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    payment_id INT REFERENCES payments(id) ON DELETE CASCADE,
    usage_gb DECIMAL(10, 2) DEFAULT 0.00,  -- Store usage in GB
    billing_cycle_start TIMESTAMP WITH TIME ZONE NOT NULL,  -- Start of the billing cycle (month)
    billing_cycle_end TIMESTAMP WITH TIME ZONE NOT NULL,    -- End of the billing cycle (month)
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMIT;