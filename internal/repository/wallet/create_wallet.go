package wallet

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/converter"
	"github.com/google/uuid"
)

func (r *repo) CreateWallet(ctx context.Context) (*model.Wallet, error) {
	wallet, err := r.DB.CreateWallet(
		ctx,
		uuid.New(),
	)

	if err != nil {
		return nil, err
	}

	return converter.ToSqlcWalletToModelWallet(&wallet), nil
}
