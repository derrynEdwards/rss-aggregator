module github.com/derrynEdwards/rss-aggregator

go 1.21.1

replace github.com/derrynEdwards/rss-aggregator/handlers v0.0.0 => ./handlers

replace github.com/derrynEdwards/rss-aggregator/internal/database v0.0.0 => ./internal/database

replace github.com/derrynEdwards/rss-aggregator/internal/utils v0.0.0 => ./internal/utils

replace github.com/derrynEdwards/rss-aggregator/internal/auth v0.0.0 => ./internal/auth

replace github.com/derrynEdwards/rss-aggregator/docs v0.0.0 => ./docs

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/go-chi/httplog v0.3.1 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	github.com/swaggo/files v1.0.1 // indirect
	github.com/swaggo/http-swagger v1.3.4 // indirect
	github.com/swaggo/swag v1.16.2 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
