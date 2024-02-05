package wallet

import (
	"context"
	"database/sql"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/converter"
)

func (r *repo) SendMoney(ctx context.Context, trans *model.Transfer) error {
	// start transaction
	tx, err := r.TX.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		return err
	}

	txDB := r.DB.WithTx(tx)

	// Check if the wallet exists
	_, err = txDB.GetWallet(ctx, converter.FromModelTranferToId(trans.From))
	if err != nil {
		tx.Rollback()
		return model.ErrWalletNotFound
	}

	// Try to debit the wallet
	_, err = txDB.DebitWallet(ctx, *converter.FromModelTransferToSqlcDebitWallet(trans))
	if err != nil {
		tx.Rollback()
		return model.ErrInsufficientFunds
	}

	// Try to credit the wallet
	_, err = txDB.CreditWallet(ctx, *converter.FromModelTransferToSqlcCreditWallet(trans))
	if err != nil {
		tx.Rollback()
		return model.ErrTargetWalletNotFound
	}

	// Write the transfer to the database
	err = txDB.CreateTransfer(ctx, *converter.FromModelTransferSqlcTransfer(trans))
	if err != nil {
		tx.Rollback()
		return model.ErrTransferFailed
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return model.ErrTransactionFailed
	}

	return nil
}
