-- name: CreateWallet :one
INSERT INTO wallets (id)
VALUES ($1)
RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE id = $1;

-- name: DebitWallet :one
UPDATE wallets
SET amount = amount - $2, 
    updated_at = NOW()
WHERE id = $1
AND amount >= $2
RETURNING *;

-- name: CreditWallet :one
UPDATE wallets
SET amount = amount + $2, 
    updated_at = NOW()
WHERE id = $1
RETURNING *;