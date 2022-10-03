#include .env
export

.PHONY:
.DEFAULT_GOAL := help

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ## Run calculator manager
	docker-compose up --remove-orphans -d --build calculator_manager

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

lint: ## Not done
	golangci-lint run

gen: ## Not done
	# mockgen -source=internal/controller/controller.go -destination=internal/controller/mocks/mock.go
	# mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go

