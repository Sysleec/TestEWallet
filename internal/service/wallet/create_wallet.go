package wallet

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
)

func (s *serv) CreateWallet(ctx context.Context) (*model.Wallet, error) {
	wlt, err := s.wRepo.CreateWallet(ctx)

	if err != nil {
		return nil, err
	}

	return wlt, err
}
