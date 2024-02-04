package wallet

import (
	"database/sql"

	"github.com/Sysleec/TestEWallet/internal/repository"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/sqlc"
)

type repo struct {
	DB *sqlc.Queries
	TX *sql.DB
}

// NewRepo creates new wallet repository
func NewRepo(db *sqlc.Queries, tx *sql.DB) repository.WalletRepository {
	return &repo{
		DB: db,
		TX: tx,
	}
}
