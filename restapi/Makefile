.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t kngnkg/tunetrail/api:${DOCKER_TAG} \
		--target deploy ./

test: ## Execute tests
	go test -race -shuffle=on ./...

DB_HOST := tunetrail-db
DB_PORT := 5432
DB_USER := tunetrail
DB_PASSWORD := tunetrail
DB_NAME := tunetrail
dry-migrate: ## Try migration
	psqldef --dry-run --host=${DB_HOST} --port=${DB_PORT} --user=${DB_USER} --password=${DB_PASSWORD} ${DB_NAME} < ../database/_tools/postgres/schema.sql

migrate:  ## Execute migration
	psqldef --host=${DB_HOST} --port=${DB_PORT} --user=${DB_USER} --password=${DB_PASSWORD} ${DB_NAME} < ../database/_tools/postgres/schema.sql

generate: ## Generate codes
	go generate ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
