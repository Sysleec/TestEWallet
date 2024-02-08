package wallet

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Sysleec/TestEWallet/internal/converter"
	"github.com/Sysleec/TestEWallet/internal/model"
	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) SendMoney(w http.ResponseWriter, r *http.Request) {
	fromWallet := chi.URLParam(r, "walletid")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	params := model.TransferReq{}
	err := decoder.Decode(&params)
	if err != nil {
		resp.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	trs := converter.FromTransRequestToModel(fromWallet, params)

	err = s.walletService.SendMoney(r.Context(), trs)

	if err != nil {
		switch {
		case errors.Is(err, model.ErrWalletNotFound):
			resp.RespondWithError(w, http.StatusNotFound, "your wallet not found")
			return
		case errors.Is(err, model.ErrInsufficientFunds):
			resp.RespondWithError(w, http.StatusBadRequest, "insufficient funds")
			return
		case errors.Is(err, model.ErrTargetWalletNotFound):
			resp.RespondWithError(w, http.StatusBadRequest, "target wallet not found")
			return
		default:
			resp.RespondWithError(w, http.StatusInternalServerError, "Failed to send money")
			return
		}
	}

	log.Info().Msgf("Money sent from wallet: %v", fromWallet)
	log.Debug().Msgf("Transfer: %v", trs)

	resp.RespondWithJSON(w, http.StatusOK, "Money sent successfully")
}
