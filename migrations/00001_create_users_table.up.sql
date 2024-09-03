BEGIN;
create table if not exists users (
    id SERIAL PRIMARY KEY,
    chat_id bigint, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
COMMIT;