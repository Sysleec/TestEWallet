package wallet

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Sysleec/TestEWallet/internal/converter"
	resp "github.com/Sysleec/TestEWallet/internal/utils"
	"github.com/go-chi/chi/v5"
)

func (s *Implementation) SendMoney(w http.ResponseWriter, r *http.Request) {
	fromWallet := chi.URLParam(r, "walletid")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	params := converter.TransferReqParams{}
	err := decoder.Decode(&params)
	if err != nil {
		resp.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	trs := converter.FromTransRequestToModel(fromWallet, params)

	err = s.walletService.SendMoney(r.Context(), trs)

	if err != nil {
		if strings.Contains(err.Error(), "your wallet not found") {
			resp.RespondWithError(w, http.StatusNotFound, "your wallet not found")
			return
		}
		if strings.Contains(err.Error(), "insufficient funds") {
			resp.RespondWithError(w, http.StatusBadRequest, "insufficient funds")
			return
		}
		if strings.Contains(err.Error(), "target wallet not found") {
			resp.RespondWithError(w, http.StatusBadRequest, "target wallet not found")
			return
		}
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp.RespondWithJSON(w, http.StatusOK, "Money sent successfully")
}
