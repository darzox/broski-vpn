BEGIN;

create table if not exists users_keys (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    access_key TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMIT;