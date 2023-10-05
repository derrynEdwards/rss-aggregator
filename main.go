package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/derrynEdwards/rss-aggregator/handlers"
	"github.com/derrynEdwards/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON:    true,
		Concise: true,
	})

	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	apiCfg := handlers.ApiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))
	apiRouter := chi.NewRouter()

	apiRouter.Get("/readiness", handlers.HandlerReadiness)
	apiRouter.Get("/err", handlers.HandlerErr)

	apiRouter.Post("/users", apiCfg.HandlerUsersCreate)
	apiRouter.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerUsersGet))

	apiRouter.Get("/feeds", apiCfg.HandlerGetFeeds)
	apiRouter.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerFeedCreate))

	apiRouter.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerFeedFollowsGet))
	apiRouter.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerFeedFollowCreate))
	apiRouter.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerFeedFollowDelete))

	router.Mount("/v1", apiRouter)

	corsMux := middleWareLog(middlewareCors(router))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
