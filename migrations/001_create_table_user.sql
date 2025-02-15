CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    coin_balance INT NOT NULL CHECK (coin_balance >= 0)
);