package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Sysleec/TestEWallet/internal/api/wallet"
	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/service/mocks"
)

func TestCreateWallet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		apiErr = fmt.Errorf("api error")

		id  = gofakeit.UUID()
		bal = gofakeit.Float64()

		res = &model.Wallet{
			ID:      id,
			Balance: bal,
		}
	)

	tests := []struct {
		name              string
		args              args
		want              *model.Wallet
		err               error
		walletServiceMock func(mc *minimock.Controller) *mocks.WalletServiceMock
		expectedCode      int
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
			},
			want: res,
			err:  nil,
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				mock := mocks.NewWalletServiceMock(mc)
				mock.CreateWalletMock.Expect(ctx).Return(res, nil)
				return mock
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
			},
			want: nil,
			err:  apiErr,
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				mock := mocks.NewWalletServiceMock(mc)
				mock.CreateWalletMock.Expect(ctx).Return(nil, apiErr)
				return mock
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			walletServiceMock := tt.walletServiceMock(mc)

			rr := httptest.NewRecorder()

			req, err := http.NewRequest("POST", "/wallet", nil)
			require.NoError(t, err)

			req = req.WithContext(tt.args.ctx)

			handler := wallet.NewImplementation(walletServiceMock)

			handler.Create(rr, req)

			if tt.err != nil {
				require.Equal(t, tt.expectedCode, rr.Code)
				require.Contains(t, rr.Body.String(), tt.err.Error())
			} else {
				require.Equal(t, tt.expectedCode, rr.Code)

				var wallet model.Wallet
				err = json.Unmarshal(rr.Body.Bytes(), &wallet)
				require.NoError(t, err)
				require.Equal(t, tt.want, &wallet)
			}
		})
	}
}
