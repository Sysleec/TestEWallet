package wallet

import (
	"github.com/Sysleec/TestEWallet/internal/repository"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/sqlc"
)

type repo struct {
	DB *sqlc.Queries
}

// NewRepo creates new wallet repository
func NewRepo(db *sqlc.Queries) repository.WalletRepository {
	return &repo{DB: db}
}
