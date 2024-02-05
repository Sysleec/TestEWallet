package wallet

import (
	"context"
	"fmt"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/converter"
)

func (r *repo) GetWallet(ctx context.Context, id string) (*model.Wallet, error) {
	wallet, err := r.DB.GetWallet(ctx, converter.FromModelTranferToId(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}

	return converter.FromSqlcWalletToModelWallet(&wallet), nil

}
