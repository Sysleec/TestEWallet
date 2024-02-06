package converter

import (
	"github.com/Sysleec/TestEWallet/internal/model"
)

func FromTransRequestToModel(from string, trs model.TransferReq) *model.Transfer {
	return &model.Transfer{
		From:   from,
		To:     trs.To,
		Amount: trs.Amount,
	}
}
