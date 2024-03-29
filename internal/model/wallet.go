package model

import "time"

type Wallet struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type Transfer struct {
	Time   time.Time `json:"time"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
}

type TransferReq struct {
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
