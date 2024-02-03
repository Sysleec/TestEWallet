-- +goose Up
CREATE TABLE IF NOT EXISTS
    wallets (
        id UUID PRIMARY KEY,
        amount FLOAT NOT NULL DEFAULT 100,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down
DROP TABLE IF EXISTS wallets;