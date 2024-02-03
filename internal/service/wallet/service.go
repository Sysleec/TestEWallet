package wallet

import "github.com/Sysleec/TestEWallet/internal/repository"

type serv struct {
	wRepo repository.WalletRepository
}

// NewService returns a new wallet service
func NewService(walletRepo repository.WalletRepository) *serv {
	return &serv{
		wRepo: walletRepo,
	}
}
