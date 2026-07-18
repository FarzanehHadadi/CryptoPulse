include .env
export
.PHONY: \
	docker-up \
	docker-down \
	migrate-up \
	migrate-down \
	migrate-status \
	generate \
	db-refresh \
	prepare \
	dev \
	run \
	test
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5434/$(DB_NAME)?sslmode=$(DB_SSLMODE)

docker-up:
	docker compose up -d postgres redis

docker-down:
	docker compose down

migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down

migrate-status:
	goose -dir migrations postgres "$(DB_URL)" status

generate:
	sqlc generate

db-refresh: 
	migrate-up generate
dev:
	air

prepare:
	docker-up migrate-up generate


run:
	go run ./cmd/api

test:
	go test ./...