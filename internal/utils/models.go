package utils

import (
	"time"

	"github.com/derrynEdwards/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// User model info
type User struct {
	ID        uuid.UUID `json:"id" example:"03f474bc-2651-4a2f-ab3e-9525d1ee40bf"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-05T00:47:49.007148Z"`
	UpdatedAt time.Time `json:"update_at" example:"2023-10-05T00:47:49.007148Z"`
	Name      string    `json:"name" example:"John"`
	ApiKey    string    `json:"api_key" example:"333a7712b7edb73922f9e5c0421fe224443247aa80da93b0409377281b38207c"`
} //@name User

func DatabaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

// Feed model info
type Feed struct {
	ID        uuid.UUID `json:"id" example:"4bda8766-37a2-46e7-ad95-1c4038b3bcfe"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-05T00:49:11.897425Z"`
	UpdatedAt time.Time `json:"update_at" example:"2023-10-05T00:49:11.897425Z"`
	Name      string    `json:"name" example:"The Example Blog"`
	Url       string    `json:"url" example:"https://example.blog"`
	UserID    uuid.UUID `json:"user_id" example:"03f474bc-2651-4a2f-ab3e-9525d1ee40bf"`
} //@name Feed

func DatabaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func DatabaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))

	for i, feed := range feeds {
		result[i] = DatabaseFeedToFeed(feed)
	}

	return result
}

// FeedFollow model info
type FeedFollow struct {
	ID        uuid.UUID `json:"id" example:"1323e287-5e9d-4999-8373-35293f38c9d8"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-05T03:57:38.673971Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-10-05T03:57:38.673971Z"`
	UserID    uuid.UUID `json:"user_id" example:"03f474bc-2651-4a2f-ab3e-9525d1ee40bf"`
	FeedID    uuid.UUID `json:"feed_id" example:"4bda8766-37a2-46e7-ad95-1c4038b3bcfe"`
} //@name FeedFollow

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))

	for i, feedFollow := range feedFollows {
		result[i] = DatabaseFeedFollowToFeedFollow(feedFollow)
	}

	return result
}
