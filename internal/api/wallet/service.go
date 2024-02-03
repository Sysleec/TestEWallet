package wallet

import "github.com/Sysleec/TestEWallet/internal/service"

// Implementation is the people service
type Implementation struct {
	walletService service.WalletService
}

// NewImplementation returns a new people service
func NewImplementation(wService service.WalletService) *Implementation {
	return &Implementation{walletService: wService}
}
