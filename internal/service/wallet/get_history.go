package wallet

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
)

func (s *serv) GetHistory(ctx context.Context, id string) ([]model.Transfer, error) {
	history, err := s.wRepo.GetHistory(ctx, id)
	if err != nil {
		return nil, err
	}

	return history, nil
}
