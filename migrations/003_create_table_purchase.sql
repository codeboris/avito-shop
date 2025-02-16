CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    merch_id INT NOT NULL REFERENCES merch(id),
    created_at TIMESTAMP NOT NULL
);