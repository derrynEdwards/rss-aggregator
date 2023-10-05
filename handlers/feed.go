package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/derrynEdwards/rss-aggregator/internal/database"
	"github.com/derrynEdwards/rss-aggregator/internal/utils"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, struct {
		Feed       utils.Feed       `json:"feed"`
		FeedFollow utils.FeedFollow `json:"feed_follow"`
	}{
		Feed:       utils.DatabaseFeedToFeed(feed),
		FeedFollow: utils.DatabaseFeedFollowToFeedFollow(feedFollow),
	})
}

func (cfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFeedsToFeeds(feeds))
}
