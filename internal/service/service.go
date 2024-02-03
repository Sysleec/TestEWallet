package service

import "github.com/Sysleec/TestEWallet/internal/model"

type WalletService interface {
	CreateWallet() (*model.Wallet, error)
	SendMoney(from, to string, amount int) error
	GetHistory(id string) ([]model.Transaction, error)
	GetWallet(id string) (*model.Wallet, error)
}
