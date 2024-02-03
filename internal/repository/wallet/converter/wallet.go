package converter

import (
	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/sqlc"
)

func ToSqlcWalletToModelWallet(wallet *sqlc.Wallet) *model.Wallet {
	return &model.Wallet{
		ID:      wallet.ID.String(),
		Balance: wallet.Amount,
	}
}
