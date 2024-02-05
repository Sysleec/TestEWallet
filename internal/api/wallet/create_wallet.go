package wallet

import (
	"net/http"

	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Create(w http.ResponseWriter, r *http.Request) {
	wallet, err := s.walletService.CreateWallet(r.Context())

	if err != nil {
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msgf("Wallet created: %v", wallet.ID)
	log.Debug().Msgf("Wallet: %v", wallet)

	resp.RespondWithJSON(w, http.StatusCreated, wallet)
}
