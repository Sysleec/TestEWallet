package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sysleec/TestEWallet/internal/api/wallet"
	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/service/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/go-chi/chi/v5"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGetHistory(t *testing.T) {
	type args struct {
		ctx      context.Context
		walletid string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		//apiErr = fmt.Errorf("api error")

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
		name              string
		args              args
		want              []model.Transfer
		err               error
		walletServiceMock func(mc *minimock.Controller) *mocks.WalletServiceMock
		expectedCode      int
	}{
		{
			name: "success case",
			args: args{
				ctx:      ctx,
				walletid: walletid,
			},
			want: res,
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				m := mocks.NewWalletServiceMock(mc)
				m.GetHistoryMock.Expect(minimock.AnyContext, walletid).Return(res, nil)
				return m
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "error case",
			args: args{
				ctx:      ctx,
				walletid: walletid,
			},
			want: nil,
			err:  model.ErrGetHistoryFailed,
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				m := mocks.NewWalletServiceMock(mc)
				m.GetHistoryMock.Expect(minimock.AnyContext, walletid).Return(nil, model.ErrGetHistoryFailed)
				return m
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			walletServiceMock := tt.walletServiceMock(mc)
			handler := wallet.NewImplementation(walletServiceMock)

			req, err := http.NewRequest("GET", "/wallet/"+tt.args.walletid+"/history", nil)
			require.NoError(t, err)

			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("walletid", tt.args.walletid)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

			rr := httptest.NewRecorder()

			handler.History(rr, req)

			if tt.err != nil {
				require.Equal(t, tt.expectedCode, rr.Code)
				require.Contains(t, rr.Body.String(), tt.err.Error())
			} else {
				require.Equal(t, tt.expectedCode, rr.Code)

				var got []model.Transfer
				err = json.Unmarshal(rr.Body.Bytes(), &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
