package wallet

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	return converter.FromSqlcWalletToModelWallet(&wallet), nil
}
