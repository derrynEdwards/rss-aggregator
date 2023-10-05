RSS Aggregator
==============

Overview
--------
Go guided project from [Boot.dev](https://boot.dev)

We're going to build an RSS feed aggregator in Go. This web server will allos clients to:
- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

Learning Goals
--------------
- Learn how to integrate a Go server with PostgreSQL
- Learn about the basics of database migrations
- Learn about long-running service workers

Requirements
------------
- Go 1.21.1
- Goose `go install github.com/pressly/goose/v3/cmd/goose@latest`
- PostgreSQL 15

Setup
-----
- Have a PostgreSQL Database Up and Running
- Change directory to `sql/schema`
- Run Goose DB Migration to setup tables
  - `goose postgres postgres://{user}:{password}@localhost:5432/rss-aggregator up`
- Install dependencies `go mod download`
- Build `go build`

Execution
---------
- Run the Built Executable `./rss-aggregator`

API Docs
--------


File Tree
---------
```shell
rss-aggregator
├── README.md
├── go.mod
├── go.sum
├── handlers
│   ├── auth.go
│   ├── err.go
│   ├── feed.go
│   ├── feed_follows.go
│   ├── readiness.go
│   └── user.go
├── headers.go
├── internal
│   ├── auth
│   │   └── auth.go
│   ├── database
│   │   ├── db.go
│   │   ├── feed_follows.sql.go
│   │   ├── feeds.sql.go
│   │   ├── models.go
│   │   └── users.sql.go
│   └── utils
│       ├── json.go
│       └── models.go
├── logger.go
├── main.go
├── rss-aggregator
└── sql
    └── schema
        ├── 001_users.sql
        ├── 002_users_apikey.sql
        ├── 003_feeds.sql
        └── 004_follows.sql
```

Sample Run
----------
```shell
❯ ./rss-aggregator
2023/10/04 22:41:19 Serving on :8080

```
