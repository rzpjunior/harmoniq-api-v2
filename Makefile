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

dev-air: $(AIR) ## Starts AIR ( Continuous Development app).
	@ air

docker-stop:
	@ docker-compose down

docker-teardown:
	@ docker-compose down --remove-orphans -v

docker-log:
	@ docker logs -f superindo_task_api

# ~~~ Cleans ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

clean: clean-artifacts clean-docker

clean-artifacts: ## Removes Artifacts (*.out)
	@ printf "Cleaning artifacts... "
	@ rm -f *.out
	@ echo "done."

clean-docker: ## Removes dangling docker images
	@ docker image prune -f
