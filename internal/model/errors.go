package model

import "errors"

var (
	ErrWalletNotFound       = errors.New("your wallet not found")
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrTargetWalletNotFound = errors.New("target wallet not found")
	ErrTransferFailed       = errors.New("transfer failed")
	ErrTransactionFailed    = errors.New("transaction failed")
)
