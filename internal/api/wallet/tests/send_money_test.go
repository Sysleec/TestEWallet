package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sysleec/TestEWallet/internal/api/wallet"
	"github.com/Sysleec/TestEWallet/internal/converter"
	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/service/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/go-chi/chi/v5"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSendMoney(t *testing.T) {
	type args struct {
		ctx      context.Context
		walletid string
		req      model.TransferReq
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		walletid = gofakeit.UUID()
		tom      = gofakeit.UUID()
		amt      = gofakeit.Float64()

		req = model.TransferReq{
			To:     tom,
			Amount: amt,
		}
	)

	tests := []struct {
		name              string
		args              args
		err               error
		walletServiceMock func(mc *minimock.Controller) *mocks.WalletServiceMock
		expectedCode      int
	}{
		{
			name: "success case",
			args: args{
				ctx:      ctx,
				walletid: walletid,
				req:      req,
			},
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				m := mocks.NewWalletServiceMock(mc)
				m.SendMoneyMock.Expect(minimock.AnyContext, converter.FromTransRequestToModel(walletid, req)).Return(nil)
				return m
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "error case - wallet not found",
			args: args{
				ctx:      ctx,
				walletid: walletid,
				req:      req,
			},
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				m := mocks.NewWalletServiceMock(mc)
				m.SendMoneyMock.Expect(minimock.AnyContext, converter.FromTransRequestToModel(walletid, req)).Return(model.ErrWalletNotFound)
				return m
			},
			expectedCode: http.StatusNotFound,
		},
		{
			name: "error case - insufficient funds",
			args: args{
				ctx:      ctx,
				walletid: walletid,
				req:      req,
			},
			walletServiceMock: func(mc *minimock.Controller) *mocks.WalletServiceMock {
				m := mocks.NewWalletServiceMock(mc)
				m.SendMoneyMock.Expect(minimock.AnyContext, converter.FromTransRequestToModel(walletid, req)).Return(model.ErrInsufficientFunds)
				return m
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			walletServiceMock := tt.walletServiceMock(mc)
			handler := wallet.NewImplementation(walletServiceMock)

			body, _ := json.Marshal(tt.args.req)
			req, err := http.NewRequest("POST", "/wallet/"+tt.args.walletid+"/send", bytes.NewBuffer(body))
			require.NoError(t, err)

			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("walletid", tt.args.walletid)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

			rr := httptest.NewRecorder()

			handler.SendMoney(rr, req)

			require.Equal(t, tt.expectedCode, rr.Code)
		})
	}
}
