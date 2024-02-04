package wallet

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/converter"
)

func (r *repo) GetHistory(ctx context.Context, id string) ([]model.Transfer, error) {
	history, err := r.DB.GetHistory(ctx, converter.FromModelTranferToId(id))
	if err != nil {
		return nil, err
	}

	return converter.FromSqlcHistorySliceToModelTransferSlice(history), nil

}
