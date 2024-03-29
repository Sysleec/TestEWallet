package wallet

import (
	"net/http"

	"github.com/Sysleec/TestEWallet/internal/model"
	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Wallet(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "walletid")

	wallet, err := s.walletService.GetWallet(r.Context(), walletID)

	if err != nil {
		log.Err(err).Msgf("Failed to get wallet: %v", walletID)
		resp.RespondWithError(w, http.StatusNotFound, model.ErrWalletNotFound.Error())
		return
	}

	log.Info().Msgf("Get wallet: %v", wallet.ID)
	log.Debug().Msgf("Wallet: %v", wallet)

	resp.RespondWithJSON(w, http.StatusOK, wallet)
}
