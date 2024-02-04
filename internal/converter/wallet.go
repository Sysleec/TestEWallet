package converter

import (
	"github.com/Sysleec/TestEWallet/internal/model"
)

type TransferReqParams struct {
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func FromTransRequestToModel(from string, trs TransferReqParams) *model.Transfer {
	return &model.Transfer{
		From:   from,
		To:     trs.To,
		Amount: trs.Amount,
	}
}
