.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image
	docker build -t lambda-test:${DOCKER_TAG} \
		--target test ./

build-no-cache: ## Build docker image without cache
	docker build --no-cache -t lambda-test:${DOCKER_TAG} \
		--target test ./

run: ## Run docker image
	docker run -p 9000:8080 \
		--rm \
		-e TUNETRAIL_DB_HOST=localhost \
		-e TUNETRAIL_DB_PORT=5555 \
		lambda-test:${DOCKER_TAG}

request:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" \
		-d '{"Name": "World"}'

test: ## Execute tests
	go test -race -shuffle=on ./...

DB_HOST := tunetrail-db
DB_PORT := 5432
DB_USER := tunetrail
DB_PASSWORD := tunetrail
DB_NAME := tunetrail
dry-migrate: ## Try migration
	psqldef --dry-run \
		--host=${DB_HOST} \
		--port=${DB_PORT} \
		--user=${DB_USER} \
		--password=${DB_PASSWORD} \
		${DB_NAME} \
		< ./_tools/postgres/schema.sql

migrate:  ## Execute migration
	psqldef \
		--host=${DB_HOST} \
		--port=${DB_PORT} \
		--user=${DB_USER} \
		--password=${DB_PASSWORD} \
		${DB_NAME} \
		< ./_tools/postgres/schema.sql

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
