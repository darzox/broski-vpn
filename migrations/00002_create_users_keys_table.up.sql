BEGIN;

create table if not exists users_keys (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    key_id bigint not null,
    access_key TEXT UNIQUE NOT NULL,
    expiration_date TIMESTAMP not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMIT;