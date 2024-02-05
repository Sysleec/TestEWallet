package wallet

import (
	"net/http"

	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Wallet(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "walletid")

	wallet, err := s.walletService.GetWallet(r.Context(), walletID)

	if err != nil {
		resp.RespondWithError(w, http.StatusNotFound, "your wallet not found")
		return
	}

	log.Info().Msgf("Get wallet: %v", wallet.ID)
	log.Debug().Msgf("Wallet: %v", wallet)

	resp.RespondWithJSON(w, http.StatusOK, wallet)
}
