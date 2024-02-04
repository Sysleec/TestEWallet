package converter

import (
	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/sqlc"
	"github.com/google/uuid"
)

func FromSqlcHistoryToModelTransfer(transfer *sqlc.History) *model.Transfer {
	return &model.Transfer{
		Amount: transfer.Amount,
		From:   transfer.FromWallet.String(),
		To:     transfer.ToWallet.String(),
		Time:   transfer.Time,
	}
}

func FromSqlcHistorySliceToModelTransferSlice(transfers []sqlc.History) []model.Transfer {
	var result []model.Transfer
	for _, transfer := range transfers {
		result = append(result, *FromSqlcHistoryToModelTransfer(&transfer))
	}
	return result
}

func FromSqlcWalletToModelWallet(wallet *sqlc.Wallet) *model.Wallet {
	return &model.Wallet{
		ID:      wallet.ID.String(),
		Balance: wallet.Amount,
	}
}

func FromModelTranferToId(walletid string) (fromWalletID uuid.UUID) {
	fromWalletID, _ = uuid.Parse(walletid)

	return fromWalletID
}
func FromModelTransferToSqlcDebitWallet(transfer *model.Transfer) *sqlc.DebitWalletParams {
	id, _ := uuid.Parse(transfer.From)
	return &sqlc.DebitWalletParams{
		ID:     id,
		Amount: transfer.Amount,
	}
}

func FromModelTransferToSqlcCreditWallet(transfer *model.Transfer) *sqlc.CreditWalletParams {
	id, _ := uuid.Parse(transfer.To)
	return &sqlc.CreditWalletParams{
		ID:     id,
		Amount: transfer.Amount,
	}
}

func FromModelTransferSqlcTransfer(transfer *model.Transfer) *sqlc.CreateTransferParams {
	fromWalletID, _ := uuid.Parse(transfer.From)
	toWalletID, _ := uuid.Parse(transfer.To)
	return &sqlc.CreateTransferParams{
		ID:         uuid.New(),
		FromWallet: fromWalletID,
		ToWallet:   toWalletID,
		Amount:     transfer.Amount,
	}
}
