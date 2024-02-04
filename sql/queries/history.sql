-- name: CreateTransfer :exec
INSERT INTO history (id, amount, from_wallet, to_wallet)
VALUES ($1, $2, $3, $4);