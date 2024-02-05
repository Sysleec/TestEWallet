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
		resp.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	log.Info().Msgf("History for wallet: %v", walletID)
	log.Debug().Msgf("History: %v", history)

	resp.RespondWithJSON(w, http.StatusOK, history)
}
