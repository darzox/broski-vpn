BEGIN;
CREATE TABLE if not exists plans (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    data_limit_gb smallint NOT NULL,
    price DECIMAL(10, 2) NOT NULL, -- Price for the plan in USD
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert the 50GB plan into the plans table
INSERT INTO plans (name, data_limit_gb, price) VALUES ('50GB Plan', 50, 10.00);


COMMIT;