CREATE TABLE purchases (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    merch_id UUID NOT NULL REFERENCES merch(id),
    purchase_date TIMESTAMP NOT NULL
);