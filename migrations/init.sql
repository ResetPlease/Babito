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

CREATE TABLE Operations 
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    amount INTEGER NOT NULL,
    target_user_id INTEGER REFERENCES Users(id) ON DELETE SET NULL,
    item VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Inventory
(
    user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES Products(id) ON DELETE CASCADE,
    count INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_username ON Users(username);
CREATE INDEX operation_idx_user_id ON Operations(user_id);
CREATE INDEX inventory_idx_user_id ON Inventory(user_id);
