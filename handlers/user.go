package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/derrynEdwards/rss-aggregator/internal/database"
	"github.com/derrynEdwards/rss-aggregator/internal/utils"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func (cfg *ApiConfig) HandlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode Parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Fatalf("%s", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseUserToUser(user))
}

func (cfg *ApiConfig) HandlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseUserToUser(user))
}
