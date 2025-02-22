CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    from_user_id INT NOT NULL REFERENCES users(id),
    to_user_id INT NOT NULL REFERENCES users(id),
    amount INT NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP NOT NULL
);