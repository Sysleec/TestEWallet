package wallet

import (
	"net/http"

	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) History(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "walletid")

	history, err := s.walletService.GetHistory(r.Context(), walletID)

	if err != nil {
		log.Err(err).Msgf("Failed to get history for wallet: %v", walletID)
		resp.RespondWithError(w, http.StatusNotFound, "wallet not found")
		return
	}

	log.Info().Msgf("History for wallet: %v", walletID)
	log.Debug().Msgf("History: %v", history)

	resp.RespondWithJSON(w, http.StatusOK, history)
}
