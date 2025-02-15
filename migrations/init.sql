-- 1k RPS выжималка
ALTER SYSTEM SET max_connections = 1000;
ALTER SYSTEM SET shared_buffers = '8GB';
ALTER SYSTEM SET effective_cache_size = '24GB';
ALTER SYSTEM SET work_mem = '64MB';
ALTER SYSTEM SET maintenance_work_mem = '1GB';
ALTER SYSTEM SET max_parallel_workers_per_gather = 4;
ALTER SYSTEM SET parallel_tuple_cost = 0.1;
ALTER SYSTEM SET parallel_setup_cost = 0.1;
ALTER SYSTEM SET wal_level = replica;
ALTER SYSTEM SET checkpoint_timeout = '10min';
ALTER SYSTEM SET synchronous_commit = off;
ALTER SYSTEM SET fsync = off;

SELECT pg_reload_conf();


CREATE TABLE Users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    hashed_password VARCHAR(128) NOT NULL,
    balance INTEGER NOT NULL CHECK (balance >= 0)
);

CREATE TABLE Products
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL
);

CREATE TYPE status_enum AS ENUM ('transfer', 'purchase');

CREATE TABLE Operations 
(
    user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
    type status_enum NOT NULL,
    amount INTEGER NOT NULL,
    target_user_id INTEGER REFERENCES Users(id) ON DELETE SET NULL,
    item VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_username ON Users(username);
CREATE INDEX operation_idx_user_id ON Operations(user_id);
CREATE INDEX operation_idx_target_user_id ON Operations(target_user_id);

INSERT INTO Products (name, price)
VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500)
;