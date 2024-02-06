package wallet

import (
	"net/http"

	"github.com/Sysleec/TestEWallet/internal/model"
	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Create(w http.ResponseWriter, r *http.Request) {
	wallet, err := s.walletService.CreateWallet(r.Context())

	if err != nil {
		log.Err(err).Msg("Failed to create wallet")
		resp.RespondWithError(w, http.StatusBadRequest, model.ErrWalletNotFound.Error())
		return
	}

	log.Info().Msgf("Wallet created: %v", wallet.ID)
	log.Debug().Msgf("Wallet: %v", wallet)

	resp.RespondWithJSON(w, http.StatusCreated, wallet)
}
