include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5434/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: docker-up
docker-up:
	docker compose up -d postgres redis

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: migrate-up
migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down

.PHONY: migrate-status
migrate-status:
	goose -dir migrations postgres "$(DB_URL)" status

.PHONY: generate
generate:
	sqlc generate

.PHONY: db-refresh
db-refresh: 
	migrate-up generate

.PHONY: dev
dev:
	air

prepare:
	docker-up migrate-up generate


.PHONY: run
run:
	go run ./cmd/api

.PHONY: test
test:
	go test ./...