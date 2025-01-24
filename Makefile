include .env

help:
	@echo "Available targets:"
	@echo "  up                  - Start containers using docker-compose"
	@echo "  down                - Stop containers using docker-compose"
	@echo "  db.connect          - Connect to local postgres database"
	@echo "  kill             	 - Kill process running on env port"
	@echo "  lint                - Run linter"
	@echo "  build               - Build application"
	@echo "  run                 - Run application in development mode"
	@echo "  start             	 - Build and run executable"
	@echo "  migrate name={name} - Create a new database migration"
	@echo "  migrate.up          - Run database migrations"
	@echo "  migrate.down        - Roll back database migrations"

up:
	@echo "Starting services..."
	docker compose up -d

down:
	@echo "Shutting down services..."
	docker compose down

db.connect:
	@echo "Connecting to local database..."
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}"

run: kill
	@echo "Starting development environment..."
	@air

kill:
	@lsof -ti tcp:${PORT} | xargs kill

lint:
	@echo "Linting project..."
	golangci-lint run

build:
	@echo "Building application for production..."
	go build -o bin/wtfemojis-api cmd/server/main.go

start: build
	@echo "Starting executable..."
	bin/wtfemojis-api

migrate:
	@echo "Create new database migration... $(name)"
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate.up:
	@echo "Migrating local database..."
	migrate -database ${DB_URL}?sslmode=disable -path internal/db/migrations up

migrate.down:
	@echo "Rolling back local database migrations..."
	migrate -database ${DB_URL}?sslmode=disable -path internal/db/migrations down