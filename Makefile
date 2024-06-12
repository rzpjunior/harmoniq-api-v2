# Database
MYSQL_USER ?= user
MYSQL_PASSWORD ?= password
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= harmoniq-dev

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

up: dev-env ##dev-air           ## Startup / Spinup Docker Compose and air
down: docker-stop               ## Stop Docker
destroy: docker-teardown clean  ## Teardown (removes volumes, tmp files, etc...)
log: docker-log  ## Teardown (removes volumes, tmp files, etc...)

dev-env: ## Bootstrap Environment (with a Docker-Compose help).
	@ docker compose -f compose.yaml up --detach

dev-air: $(AIR) ## Starts AIR (Continuous Development app).
	@ air

docker-stop:
	@ docker-compose down

docker-teardown:
	@ docker-compose down --remove-orphans -v

docker-log:
	@ docker logs -f superindo_task_api

# ~~~ Build and Clean Commands ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

build: clean-modcache update-deps build-api ## Build the Go application

clean-modcache: ## Clean Go module cache
	@ go clean -modcache

update-deps: ## Update dependencies and tidy up
	@ go mod tidy

build-api: ## Build the Go application for Linux
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./api

clean: clean-artifacts clean-docker

clean-artifacts: ## Removes Artifacts (*.out)
	@ printf "Cleaning artifacts... "
	@ rm -f *.out
	@ echo "done."

clean-docker: ## Removes dangling docker images
	@ docker image prune -f

# ~~~ Docker Build and Run Commands ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

docker-build: ## Build Docker image
	@ docker build -t harmoniq-api-v2 .

docker-run: ## Run Docker container
	@ docker run --rm -it -p 8080:8080 harmoniq-api-v2

# ~~~ Help ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

help: ## Display this help
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
