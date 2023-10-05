package handlers

import (
	"encoding/json"
	"github.com/derrynEdwards/rss-aggregator/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"

	"github.com/derrynEdwards/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *ApiConfig) HandlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *ApiConfig) HandlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
