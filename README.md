# Go Store: final-project-batch1-team-gilang

## Project layout
```
.
├── cmd
│   └── main.go
├── config
│   ├── app.env
│   └── app.env.example
├── docker-compose.yml
├── Dockerfile
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── controller
│   │   ├── model
│   │   ├── repository
│   │   ├── schema
│   │   └── service
│   └── pkg
│       ├── config
│       │   └── config.go
│       ├── db
│       │   └── db.go
│       └── middleware
│           ├── logging_middleware.go
│           └── recovery_middleware.go
├── Makefile
└── README.md
```

## Prerequisites

- Install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

## Running Locally

1. `cp config/app.env.sample config/app.env`, adjust the value
2. `make run` or `make run-docker`
3. App running!

```bash
❯ make help
migrate-all                    Rollback migrations, all migrations
migrate-create                 Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-down                   Rollback migrations, latest migration (1)
migrate-up                     Run migrations UP
run-docker                     Set up all environments and run the application on Docker.
run                            Running application without Docker
```
