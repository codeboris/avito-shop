CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    from_user_id UUID NOT NULL REFERENCES users(id),
    to_user_id UUID NOT NULL REFERENCES users(id),
    amount INT NOT NULL CHECK (amount > 0),
    transaction_date TIMESTAMP NOT NULL
);