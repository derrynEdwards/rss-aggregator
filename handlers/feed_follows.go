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

type FeedFollowCreateParams struct {
	FeedID uuid.UUID `json:"feed_id" example:"26296df9-59ce-4fbe-8c7d-f3397fcb91ce"`
} //@name FeedFollowCreateParams

// HandlerFeedFollowsGet godoc
// @Summary Gets all FeedFollows for authenticated user
// @Description Get all FeedFollows for authenticated user
// @Tags feed_follows
// @Produce json
// @Success 200 {array} utils.FeedFollow
// @Failure 500
// @Router /feed_follows [get]
// @Security ApiKeyAuth
func (cfg *ApiConfig) HandlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

// HandlerFeedFollowCreate godoc
// @Summary Creates a FeedFollow for authenticated user
// @Description Creates a new FeedFollow for authenticated user
// @Tags feed_follows
// @Accept json
// @Produce json
// @Param feedfollow body FeedFollowCreateParams true "feed id"
// @Success 200 {object} utils.FeedFollow
// @Failure 500
// @Router /feed_follows [post]
// @Security ApiKeyAuth
func (cfg *ApiConfig) HandlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := FeedFollowCreateParams{}
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

// HandlerFeedFollowDelete godoc
// @Summary Deletes specific FeedFollow for authenticated user
// @Description Deletes specific FeedFollow for authenticated user
// @Tags feed_follows
// @Param feedFollowID path string true "feed uuid"
// @Success 200
// @Failure 500
// @Router /feed_follows/{feedFollowID} [delete]
// @Security ApiKeyAuth
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
