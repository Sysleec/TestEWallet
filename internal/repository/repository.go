package repository

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context) (*model.Wallet, error)
	// SendMoney(ctx context.Context, trx model.Transfer) error
	// GetHistory(ctx context.Context, id string) ([]model.Transfer, error)
	// GetWallet(ctx context.Context, id string) (*model.Wallet, error)
}
