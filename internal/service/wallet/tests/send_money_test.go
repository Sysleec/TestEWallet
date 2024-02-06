package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/mocks"
	"github.com/Sysleec/TestEWallet/internal/service/wallet"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestServ_SendMoney(t *testing.T) {
	type fields struct {
		wRepo *mocks.WalletRepositoryMock
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		walletid = gofakeit.UUID()
		tom      = gofakeit.UUID()
		amt      = gofakeit.Float64()
		tm       = gofakeit.Date()

		req = &model.Transfer{
			Time:   tm,
			From:   walletid,
			To:     tom,
			Amount: amt,
		}
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success case",
			fields: fields{
				wRepo: func() *mocks.WalletRepositoryMock {
					mock := mocks.NewWalletRepositoryMock(mc)
					mock.SendMoneyMock.Expect(ctx, req).Return(nil)
					return mock
				}(),
			},
			wantErr: false,
		},
		{
			name: "error case",
			fields: fields{
				wRepo: func() *mocks.WalletRepositoryMock {
					mock := mocks.NewWalletRepositoryMock(mc)
					mock.SendMoneyMock.Expect(ctx, req).Return(errors.New("error"))
					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := wallet.NewService(tt.fields.wRepo)

			err := s.SendMoney(ctx, req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
