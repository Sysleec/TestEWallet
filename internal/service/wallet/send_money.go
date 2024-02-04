package wallet

import (
	"context"

	"github.com/Sysleec/TestEWallet/internal/model"
)

func (s *serv) SendMoney(ctx context.Context, trans *model.Transfer) error {
	err := s.wRepo.SendMoney(ctx, trans)

	if err != nil {
		return err
	}

	return nil
}
