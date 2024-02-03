package repository

import "github.com/Sysleec/TestEWallet/internal/model"

//import "github.com/Sysleec/TestEWallet/internal/repository/model"

type WalletService interface {
	CreateWallet() (*model.Wallet, error)
	SendMoney(from, to string, amount int) error
	GetHistory(id string) ([]model.Transaction, error)
	GetWallet(id string) (*model.Wallet, error)
}
