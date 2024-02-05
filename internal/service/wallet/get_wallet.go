package wallet

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
)

func (s *serv) GetWallet(ctx context.Context, id string) (*model.Wallet, error) {
	wallet, err := s.wRepo.GetWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	return wallet, nil

}
