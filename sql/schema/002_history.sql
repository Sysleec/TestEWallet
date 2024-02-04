-- +goose Up
CREATE TABLE IF NOT EXISTS
    history (
        id UUID PRIMARY KEY,
        amount FLOAT NOT NULL,
        from_wallet UUID NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,
        to_wallet UUID NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,
        time TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down
DROP TABLE IF EXISTS history;