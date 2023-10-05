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

type UserCreateParams struct {
	Name string `json:"name" example:"John"`
} //@name UserCreateParams

// HandlerUsersCreate godoc
// @Summary Creates a user
// @Description Creates a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserCreateParams true "user name"
// @Success 200 {object} utils.User
// @Failure 500
// @Router /users [post]
func (cfg *ApiConfig) HandlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := UserCreateParams{}
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

// HandlerUsersGet godoc
// @Summary Gets current user information
// @Description Gets current user information
// @Tags users
// @Produce json
// @Success 200 {object} utils.User
// @Failure 500
// @Router /users [get]
// @Security ApiKeyAuth
func (cfg *ApiConfig) HandlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseUserToUser(user))
}
