package wallet

import (
	"context"
	"database/sql"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/Sysleec/TestEWallet/internal/repository/wallet/converter"
	"github.com/rs/zerolog/log"
)

func (r *repo) SendMoney(ctx context.Context, trans *model.Transfer) (err error) {
	// start transaction
	tx, err := r.TX.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	// rollback transaction if any error occurs
	defer func() {
		if err != nil {
			log.Err(err).Msg("Failed to send money")
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Err(rbErr).Msg("Failed to rollback transaction")
			}
		}
	}()

	txDB := r.DB.WithTx(tx)
	// Check if the wallet exists
	_, err = txDB.GetWallet(ctx, converter.FromModelTranferToId(trans.From))
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ErrWalletNotFound
		}
		return err
	}

	// Try to debit the wallet
	_, err = txDB.DebitWallet(ctx, *converter.FromModelTransferToSqlcDebitWallet(trans))
	if err != nil {
		return model.ErrInsufficientFunds
	}

	// Try to credit the wallet
	_, err = txDB.CreditWallet(ctx, *converter.FromModelTransferToSqlcCreditWallet(trans))
	if err != nil {
		return model.ErrTargetWalletNotFound
	}

	// Write the transfer to the database
	err = txDB.CreateTransfer(ctx, *converter.FromModelTransferSqlcTransfer(trans))
	if err != nil {
		return model.ErrTransferFailed
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return model.ErrTransactionFailed
	}

	return nil
}
