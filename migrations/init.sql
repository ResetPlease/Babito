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