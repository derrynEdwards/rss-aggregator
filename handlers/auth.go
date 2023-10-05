package handlers

import (
	"net/http"

	"github.com/derrynEdwards/rss-aggregator/internal/auth"
	"github.com/derrynEdwards/rss-aggregator/internal/database"
	"github.com/derrynEdwards/rss-aggregator/internal/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Couldn't find an api key")
			return
		}

		user, err := cfg.DB.GetUsersByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
