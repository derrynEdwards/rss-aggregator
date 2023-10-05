package handlers

import (
	"github.com/derrynEdwards/rss-aggregator/internal/utils"
	"net/http"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
