.PHONY: help setup test build dev all clean lint run help
.PHONY: docker-build docker-run docker-kill


VERSION := 'joi-energy-golang/endpoints/standard.GitTag=$(shell git describe --tags --always)'
COMMIT := 'joi-energy-golang/endpoints/standard.GitCommit=$(shell git rev-list --oneline -1 HEAD)'
BUILD := 'joi-energy-golang/endpoints/standard.BuildData=$(shell date +%F%t%T)'

BUILD_DIR := bin
TOOLS_DIR := tools

help: ## Show help.
	@printf "A set of development commands.\n"
	@printf "\nUsage:\n"
	@printf "\t make \033[36m<commands>\033[0m\n"
	@printf "\nThe Commands are:\n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Set up the server.
	@go mod download

test: ## Run the tests.
	@go test -v ./... -cover -coverprofile=coverage.txt -covermode=atomic -p 8

build: ## Build the server.
	@go build -v -ldflags "-s -w -X $(COMMIT) -X $(VERSION) -X $(BUILD)" -o ./bin/server ./cmd/server

dev: ## Run local server.
	@go run -ldflags "-X $(COMMIT) -X $(VERSION) -X $(BUILD)" ./cmd/server/main.go

all: clean ## Run all tests, then build and run.
	@$(MAKE) lint
	@$(MAKE) test
	@$(MAKE) build
	@$(MAKE) run

clean: ## Clean up, i.e. remove build artifacts.
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)
	@go mod tidy

run: build ## Run the binary.
	$(BUILD_DIR)/server

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

lint: $(TOOLS_DIR)/golangci-lint/golangci-lint ## Run linters.
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./... --enable paralleltest

docker-build: ## Build container's Docker.
	@docker build -t server .

docker-run: ## Run container's Docker.
	@docker run --name new-server -p 8080:8080 -it server

docker-kill: ## Kill container's Docker.
	@docker kill new-server
