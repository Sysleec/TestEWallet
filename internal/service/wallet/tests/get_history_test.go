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

func TestServ_GetHistory(t *testing.T) {
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

		res = []model.Transfer{
			{
				Time:   tm,
				From:   walletid,
				To:     tom,
				Amount: amt,
			},
		}
	)

	tests := []struct {
		name    string
		fields  fields
		want    []model.Transfer
		wantErr bool
	}{
		{
			name: "success case",
			fields: fields{
				wRepo: func() *mocks.WalletRepositoryMock {
					mock := mocks.NewWalletRepositoryMock(mc)
					mock.GetHistoryMock.Expect(ctx, walletid).Return(res, nil)
					return mock
				}(),
			},
			want:    res,
			wantErr: false,
		},
		{
			name: "error case",
			fields: fields{
				wRepo: func() *mocks.WalletRepositoryMock {
					mock := mocks.NewWalletRepositoryMock(mc)
					mock.GetHistoryMock.Expect(ctx, walletid).Return(nil, errors.New("error"))
					return mock
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := wallet.NewService(tt.fields.wRepo)

			got, err := s.GetHistory(ctx, walletid)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
