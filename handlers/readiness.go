package handlers

import (
	"github.com/derrynEdwards/rss-aggregator/internal/utils"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}

	utils.RespondWithJSON(w, http.StatusOK, response{
		Status: "OK",
	})
}
