ifneq (,$(wildcard ./.env))
    include .env
    export
endif

GOOSE_DBSTRING=postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable

run:
	docker compose up --build

migrateup:
	GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose up

migratedown:
	GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose up