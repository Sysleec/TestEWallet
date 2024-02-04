-- name: CreateTransfer :exec
INSERT INTO history (id, amount, from_wallet, to_wallet)
VALUES ($1, $2, $3, $4);

-- name: GetHistory :many
SELECT * FROM history
WHERE from_wallet = $1 OR to_wallet = $1;