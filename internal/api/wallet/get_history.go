package wallet

import (
	"net/http"

	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/go-chi/chi/v5"
)

func (s *Implementation) History(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "walletid")

	wallet, err := s.walletService.GetHistory(r.Context(), walletID)

	if err != nil {
		resp.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	resp.RespondWithJSON(w, http.StatusOK, wallet)
}
