package wallet

import (
	"net/http"

	resp "github.com/Sysleec/TestEWallet/internal/utils"
)

func (s *Implementation) Create(w http.ResponseWriter, r *http.Request) {
	wallet, err := s.walletService.CreateWallet(r.Context())

	if err != nil {
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp.RespondWithJSON(w, http.StatusCreated, wallet)
}
