BEGIN;
create table if not exists users (
    id SERIAL PRIMARY KEY,
    chat_id bigint unique not null, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
COMMIT;